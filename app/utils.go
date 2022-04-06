package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"net/http"
	"os"
	"time"
)

// 输出错误信息并退出
func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

// 计算时间差，单位纳秒
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}

// 通过Http状态码判断是否请求成功 [200.500] 则认为成功 超时、大于500则认为错误
func HttpStatusCode(response *http.Response) (code int, isSucceed bool) {

	defer response.Body.Close()
	code = response.StatusCode
	if code >= http.StatusOK && code <= http.StatusInternalServerError {
		isSucceed = true
	}

	return
}

//根据域名解析IP地址
func NsLookupDomain(domain string) []string {
	var ips []string
	iprecords, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("域名解析失败 %v \n", err)
		return ips
	}
	for _, ip := range iprecords {
		ips = append(ips, ip.String())
	}
	return ips
}
