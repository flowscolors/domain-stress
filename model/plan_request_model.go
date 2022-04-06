package model

import (
	"fmt"
	"os"
)

// 请求计划
type PlanRequest struct {
	Id      string // 消息Id
	ChanId  uint64 // 消息通道Id
	Number  uint64 // 请求时间 纳秒
	Request *Request
}

// NewPlanRequest
// url 压测的url
// form 请求的格式 http/webscoket/grpc
// method 请求的方法
// timeout 请求超时时间 单位ms
func NewPlanRequest(chanId uint64, number uint64, request *Request) (planRequest *PlanRequest, err error) {

	planRequest = &PlanRequest{
		ChanId:  chanId,
		Number:  number,
		Request: request,
	}

	return

}

func PrintExecutionPlans(planRequests []PlanRequest) {
	fmt.Println("************************  执行计划 stat  ***************************")
	for _, planRequest := range planRequests {
		fmt.Fprintf(os.Stderr, "协程id %d  ，计划往目标 URL %s ，发送 %s 请求数 %d\n", planRequest.ChanId, planRequest.Request.Url, planRequest.Request.Method, planRequest.Number)
	}
	fmt.Println("************************  执行计划 end  ***************************")
}
