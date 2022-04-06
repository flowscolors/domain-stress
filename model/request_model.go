package model

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// 待发送请求
type Request struct {
	Url     string            // Url
	Form    string            // http/webSocket/grpc
	Path    string            // Path
	Method  string            // 方法 GET/POST/PUT/DELETE
	Headers map[string]string // Headers
	Body    string            // body
	Timeout time.Duration     // 请求超时时间

	// 连接以后初始化事件
	// 循环事件 切片 时间 动作
}

func (r *Request) GetBody() (body io.Reader) {
	body = strings.NewReader(r.Body)

	return
}

// NewRequest
// url 压测的url
// form 请求的格式 http/webscoket/grpc
// method 请求的方法
// timeout 请求超时时间 单位ms
func NewRequest(url string, form string, method string, path string, reqTimeout uint64, reqHeaders []string, reqBody string) (request *Request, err error) {

	var headers = make(map[string]string)

	if method == "POST" && reqBody != "" {
		headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
		for _, v := range reqHeaders {
			getHeaderValue(v, headers)
		}
	}

	//var timeout time.Duration= 30 * time.Second
	var timeout time.Duration = time.Duration(reqTimeout) * time.Microsecond

	request = &Request{
		Url:     url,
		Form:    form,
		Method:  strings.ToUpper(method),
		Path:    path,
		Headers: headers,
		Body:    reqBody,
		//Timeout: time.Duration(reqTimeout) * time.Second,
		Timeout: timeout,
	}

	return

}

func (r *Request) SetURL(form string, ip string, path string) {
	//	url := "http://www.baidu.com"
	url := form + "://" + ip + path
	r.Url = url
}

func getHeaderValue(v string, headers map[string]string) {
	index := strings.Index(v, ":")
	if index < 0 {
		return
	}

	vIndex := index + 1
	if len(v) >= vIndex {
		value := strings.TrimPrefix(v[vIndex:], " ")

		if _, ok := headers[v[:index]]; ok {
			headers[v[:index]] = fmt.Sprintf("%s; %s", headers[v[:index]], value)
		} else {
			headers[v[:index]] = value
		}
	}
}

// 打印请求
func (r *Request) Print() {
	if r == nil {

		return
	}

	result := fmt.Sprintf("request:\n form:%s \n url:%s \n method:%s \n headers:%v \n", r.Form, r.Url, r.Method, r.Headers)
	result = fmt.Sprintf("%s data:%v \n", result, r.Body)
	fmt.Println(result)

	return
}

// 请求结果
type RequestResults struct {
	Id            string // 消息Id
	ChanId        uint64 // 消息通道Id
	Time          uint64 // 请求时间 纳秒
	IsSucceed     bool   // 是否请求成功
	ReturnCode    int    // 返回码
	ReceivedBytes int64
}

func (r *RequestResults) SetId(chanId uint64, number uint64) {
	id := fmt.Sprintf("%d_%d", chanId, number)

	r.Id = id
	r.ChanId = chanId
}
