package plan

import (
	"domain-stress/model"
	"testing"
)

func TestExecutionAverage(t *testing.T) {
	ips := []string{"216.113.179.53", "64.4.253.77", "66.211.175.229", "66.211.172.37", "216.113.181.253", "209.140.148.143"}
	concurrency := 2
	totalNumber := 100
	request := &model.Request{
		Path:   "/",
		Form:   "http",
		Method: "GET",
	}
	planRequests := ExecutionAverage(ips, uint64(concurrency), uint64(totalNumber), request)
	model.PrintExecutionPlans(planRequests)
	if planRequests == nil {
		t.Error(`ExecutionAverage has error `) // 请求错误
	}

}
