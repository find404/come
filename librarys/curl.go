package librarys

/**
 * curl类
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月18日 00:30:00
 * @version     0.0.1 版本号
 */
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
	Scheme             string
	Host               string
	Mothod             string
	Url                string
	ResponseStatusCode int
	HeadersMap         map[string]string
	RequestJsonBytes   io.Reader
}

// 发起请求
func (hr HttpRequest) GetResponse(bodyStruct interface{}) {
	request, err := http.NewRequest(hr.Mothod, hr.Url, hr.RequestJsonBytes)
	request.Host = hr.HeadersMap["host"]
	for key, header := range hr.HeadersMap {
		request.Header.Set(key, header)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do err : %v", err)
		return
	}

	defer response.Body.Close()

	if hr.ResponseStatusCode == 0 {
		hr.ResponseStatusCode = ClHttpCode200
	}

	if response.StatusCode == hr.ResponseStatusCode {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			json.Unmarshal(body, &bodyStruct)
		}
	}
}

//设置get请求
func (hr *HttpRequest) SetMothodGet() {
	hr.Mothod = "GET"
}

//设置post请求
func (hr *HttpRequest) SetMothodPost() {
	hr.Mothod = "POST"
}

//设置状态码请求
func (hr *HttpRequest) SetResponseStatusCode200() {
	hr.ResponseStatusCode = ClHttpCode200
}

//设置状态码请求
func (hr *HttpRequest) SetResponseStatusCode302() {
	hr.ResponseStatusCode = ClHttpCode302
}
