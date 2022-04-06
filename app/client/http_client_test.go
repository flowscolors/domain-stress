package client

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {
	var (
		url     string
		method  string
		timeout time.Duration
	)

	method = "GET"
	url = "http://www.baidu.com"
	timeout = 5 * time.Second
	resp, requestTime, err := HttpRequest(method, url, nil, nil, timeout)
	fmt.Printf("Here is the resp  StatusCode %d and the time is %d ns: \n", resp.StatusCode, requestTime)
	if err != nil {
		t.Error(`http request is ok`) // 请求错误
	}

	method = "GET"
	url = "https://111.230.87.135:10081/api/v1/namespaces/default/services"
	headers := make(map[string]string)
	headers["Authorization"] = "Beaer eyJhbGciOiJSUzI1NiIsImtpZCI6ImFlU3h1YjEwSTZ6dThScjBOX1JHX0RFVk5EWnMwVmlSanZjdWN1NTU1MXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4tbjV4NGgiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImMwZGNiNjQxLWRiZTMtNGRkYy05N2YwLTJkYWYwYzRhYTdhZiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.dk71CfpjtAPZCu9tzvzJ7cqIcGgQkZsujHlJczoPK53YbxrMyA3wYkq-6tbiedXjuQ-zUVFh9fU6kDN-4lAxyHflfqly3-qsZ3LUyDra1WFNqM-UyQjYG8OqCQo0X9-wLTqOGO2TeF-Fz-P5_SFecIRYff6v2ZwbKwPmkK1CRRI3QSAba7Bn0uh5bU7PRyc81xlRwUY4EiI-bgGZ68D8-d3ZMyX6_mRdorRa7l7xjXjykqKDN5o6EUHWnZohVTRkFsIYjhNwmt7pk8XW0VdX0BY5yzdBNnM2SbeTY6HGDiZjSP2YKfVoRX6wvFpMd0rgHRcFJXolgUT4_yYsrQ-BBw"
	resp, requestTime, err = HttpRequest(method, url, nil, headers, timeout)
	fmt.Printf("Here is the resp  StatusCode %d and the time is %d ns: \n", resp.StatusCode, requestTime)
	if err != nil {
		t.Error(`http request is ok`) // 请求错误
	}

	method = "POST"
	url = "https://111.230.87.135:10081/api/v1/namespaces/default/services"
	body := "{\n  \"kind\": \"Service\", \n  \"apiVersion\": \"v1\", \n  \"metadata\": {\n    \"name\": \"nginx\", \n    \"labels\": {\n      \"app\": \"nginx\"\n    }, \n    \"namespace\": \"default\"\n  }, \n  \"spec\": {\n    \"selector\": {\n      \"app\": \"nginx\"\n    }, \n    \"type\": \"NodePort\", \n    \"ports\": [\n      {\n        \"name\": \"nginx\", \n        \"nodePort\": 30000, \n        \"port\": 80, \n        \"protocol\": \"TCP\", \n        \"targetPort\": 80\n      }\n    ]\n  }\n}"
	resp, requestTime, err = HttpRequest(method, url, strings.NewReader(body), headers, timeout)
	fmt.Printf("Here is the resp  StatusCode %d and the time is %d ns: \n", resp.StatusCode, requestTime)
	if err != nil {
		t.Error(`http request is ok`) // 请求错误
	}
}
