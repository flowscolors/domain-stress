package link

import (
	"domain-stress/app"
	"domain-stress/app/client"
	"domain-stress/model"
	"fmt"
	"sync"
)

// http go link 执行
func Http(chanId uint64, ch chan<- *model.RequestResults, retry uint, totalNumber uint64, wg *sync.WaitGroup, request *model.Request) {

	defer func() {
		wg.Done()
	}()

	//fmt.Printf("启动协程 编号:%05d \n", chanId)
	for i := uint64(0); i < totalNumber; i++ {

		isSucceed, returnCode, requestTime, contentLength := send(request, retry)

		requestResults := &model.RequestResults{
			Time:          requestTime,
			IsSucceed:     isSucceed,
			ReturnCode:    returnCode,
			ReceivedBytes: contentLength,
		}

		requestResults.SetId(chanId, i)

		ch <- requestResults
	}

	return
}

// send 往目标request发送一次请求 带重试
func send(request *model.Request, retry uint) (bool, int, uint64, int64) {
	var (
		isSucceed        = false
		returnCode       int
		contentLength    = int64(0)
		totalRequestTime uint64
	)
	for i := uint(1); i <= retry; i++ {
		resp, requestTime, err := client.HttpRequest(request.Method, request.Url, request.GetBody(), request.Headers, request.Timeout)
		if err == nil {
			totalRequestTime = totalRequestTime + requestTime
			// 验证请求是否成功
			returnCode, isSucceed = app.HttpStatusCode(resp)
			contentLength = resp.ContentLength
			return isSucceed, returnCode, totalRequestTime, contentLength
		} else {
			totalRequestTime = totalRequestTime + requestTime
			returnCode = 509 // 请求错误
			fmt.Println("第 ", i, " 次请求失败:", err)
		}
	}

	return isSucceed, returnCode, totalRequestTime, contentLength
}
