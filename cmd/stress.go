package cmd

import (
	"domain-stress/app"
	"domain-stress/app/link"
	"domain-stress/app/plan"
	"domain-stress/app/statistics"
	"domain-stress/model"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sync"
	"time"
)

const (
	LoadBalanceAverage = "Average"
	LoadBalanceDelay   = "Delay"

	FormTypeHttp      = "http"
	FormTypeWebSocket = "webSocket"
	FormTypeGRPC      = "grpc"
)

var (
	concurrency     uint64         = 1                  // 并发数(对于每个IP生效)
	totalNumber     uint64         = 1                  // 请求数(发压总数)
	timeout         uint64         = 5000               // 超时时间 单位 ms  构建request
	retry           uint           = 3                  // 重试次数
	loadBalanceMode                = LoadBalanceAverage // 负载均衡模式
	domainName      string                              // 压测的域名 构建request
	path            string                              // 请求路径  构建request
	method          = "GET"                             // 请求方法  构架request
	form            = FormTypeHttp                      // 请求格式  构建request
	headers         []string                            // 自定义头信息传递给服务器  构建request
	body            string                              // HTTP POST方式传送数据  构建request
)

// stressCmd represents the domain
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress all ip after domain ",
	Run: func(cmd *cobra.Command, args []string) {
		/*		fmt.Fprintf(os.Stderr, "echo , actual: %d \n", concurrency)
				fmt.Fprintf(os.Stderr, "echo , actual: %d \n", totalNumber)
				fmt.Fprintf(os.Stderr, "echo , actual: %d \n", timeout)
				fmt.Fprintf(os.Stderr, "echo , actual: %d \n", retry)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", loadBalanceMode)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", domainName)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", path)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", method)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", form)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", headers)
				fmt.Fprintf(os.Stderr, "echo , actual: %s \n", body)*/

		// 设置接收数据缓存
		ch := make(chan *model.RequestResults, 1000)
		var (
			wg          sync.WaitGroup // 发送数据完成
			wgReceiving sync.WaitGroup // 数据处理完成

		)
		wgReceiving.Add(1)

		//---------------------------------------------------------------------------

		//1.参数校验
		if concurrency == 0 || totalNumber == 0 {
			fmt.Printf("示例: ./domain_stress stress -d www.baidu.com -c 10 -n 100 \n")
			fmt.Printf("concurrency 并发数 或 totalNUmber 不能为0 \n")
			return
		}

		if domainName == "" {
			fmt.Printf("示例: ./domain_stress stress -d www.baidu.com -c 10 -n 100 \n")
			fmt.Printf("Domian域名解析必填 \n")
			return
		}
		fmt.Fprintf(os.Stderr, "向目标域名 %s 以每个后端IP %d 个协程的方式发送 %d 个请求。每个请求超时时间为 %d ms，重试次数为 %d 次。\n\n", domainName, concurrency, totalNumber, timeout, retry)

		//2.构建request
		requestModel, err := model.NewRequest("", form, method, path, timeout, headers, body)
		if err != nil {
			fmt.Printf("参数不合法 %v \n", err)
			return
		}

		//3.解析后端IP池
		ips := app.NsLookupDomain(domainName)

		//4.计算发压方案 后端6个IP，则默认起6个协程进行请求（根据参数也可以是每个IP起多个协程），每个协程发多少次，输出发压方案
		planRequests := ExecutionPlan(ips, concurrency, totalNumber, loadBalanceMode, requestModel)
		model.PrintExecutionPlans(planRequests)

		//5.执行发压  启动协程
		for _, planRequest := range planRequests {
			wg.Add(1)
			go link.Http(planRequest.ChanId, ch, retry, planRequest.Number, &wg, planRequest.Request)
		}

		//6.输出发压结果 启动计算输出携程 是否需要先与计算协程启动
		go statistics.ReceivingResults(concurrency*uint64(len(ips)), ch, &wgReceiving)

		//---------------------------------------------------------------------------

		// 等待所有的数据都发送完成
		wg.Wait()

		// 延时1毫秒 确保数据都处理完成了
		time.Sleep(1 * time.Millisecond)
		close(ch)

		// 数据全部处理完成了
		wgReceiving.Wait()

		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)

	// 解析参数
	stressCmd.PersistentFlags().Uint64VarP(&concurrency, "concurrency", "c", 1, "并发数")
	stressCmd.PersistentFlags().Uint64VarP(&totalNumber, "total", "n", 1, "请求数(单个并发/协程)")
	stressCmd.PersistentFlags().Uint64VarP(&timeout, "timeout", "t", 5000, "超时时间 单位ms")
	stressCmd.PersistentFlags().UintVarP(&retry, "retry", "r", 3, "重试次数")
	stressCmd.PersistentFlags().StringVarP(&loadBalanceMode, "mode", "M", LoadBalanceAverage, "负载均衡模式")
	stressCmd.PersistentFlags().StringVarP(&domainName, "domain", "d", "", "目标域名")
	stressCmd.PersistentFlags().StringVarP(&path, "path", "p", "/", "请求路径")
	stressCmd.PersistentFlags().StringVarP(&method, "method", "m", "GET", "请求方法")
	stressCmd.PersistentFlags().StringVarP(&form, "form", "f", FormTypeHttp, "请求格式")
	stressCmd.PersistentFlags().StringArrayVarP(&headers, "headers", "H", nil, "自定义头信息传递给服务器 示例:-H 'Content-Type: application/json'")
	stressCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "HTTP POST方式传送数据body消息体")

}

func ExecutionPlan(ips []string, concurrency uint64, totalNumber uint64, loadBalanceMode string, request *model.Request) []model.PlanRequest {
	var planRequests []model.PlanRequest

	switch loadBalanceMode {
	case LoadBalanceAverage:
		planRequests = plan.ExecutionAverage(ips, concurrency, totalNumber, request)
	case LoadBalanceDelay:
		planRequests = plan.ExecutionDelay(ips, concurrency, totalNumber, request)
	default:
		data := fmt.Sprintf("不支持的负载类型:%d", loadBalanceMode)
		panic(data)
	}

	return planRequests
}
