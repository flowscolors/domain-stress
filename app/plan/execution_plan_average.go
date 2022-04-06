package plan

import (
	"domain-stress/model"
	"fmt"
)

//按IP数和并发数 均分请求
func ExecutionAverage(ips []string, concurrency uint64, totalNumber uint64, request *model.Request) []model.PlanRequest {
	var (
		planRequests []model.PlanRequest
		chanid       uint64
	)
	totalChanNumber := int(concurrency) * len(ips)
	for i := 0; i < len(ips); i++ {
		realRequest, _ := model.NewRequest(request.Url, request.Form, request.Method, request.Path, uint64(request.Timeout), nil, request.Body)
		realRequest.SetURL(request.Form, ips[i], request.Path)
		realRequest.Headers = request.Headers

		for j := 0; j < int(concurrency); j++ {
			chanNumber := int(totalNumber) / totalChanNumber
			//遇到总数除不尽的情况是平均优先 还是总数优先 ,目前是按总数优先
			if i == len(ips)-1 && j == int(concurrency)-1 {
				chanNumber = int(totalNumber) - (chanNumber * int(chanid))
			}
			chanid++
			request.SetURL(request.Form, ips[i], request.Path)
			planRequest, err := model.NewPlanRequest(chanid, uint64(chanNumber), realRequest)
			if err != nil {
				fmt.Printf("创建执行计划错误  %v \n", err)
			}
			planRequests = append(planRequests, *planRequest)
		}
	}
	return planRequests
}
