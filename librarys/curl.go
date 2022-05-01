package librarys

/**
 * curl类
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月18日 00:30:00
 * @version     0.0.1 版本号
 */
import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
)

var httpClient = http.Client{Transport: &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}}

// GetResponseInfo 发起请求，返回repson
func GetResponseInfo(mothod, URL string, header map[string][]string, requestBytes io.Reader) (*http.Response, error) {
	var result *http.Response
	var infoError error
	if mothod == ClEmpty {
		mothod = http.MethodGet
	}
	if URL == ClEmpty {
		infoError = errors.New("URL is empty")
		return result, infoError
	}
	if header == nil {
		infoError = errors.New("header is empty")
		return result, infoError
	}
	request, newRequestError := http.NewRequest(mothod, URL, requestBytes)
	if newRequestError != nil {
		infoError = newRequestError
		return result, infoError
	}
	request.Header = header
	response, clientDoError := httpClient.Do(request)
	if clientDoError != nil {
		infoError = clientDoError
		return result, infoError
	}

	result = response
	infoError = nil
	return result, infoError
}

// GetBody 发起请求，返回body
func GetBody(mothod, URL string, header map[string][]string, requestBytes io.Reader) ([]byte, error) {
	var bodyByte []byte
	var bodyError error
	responseInfo, responseInfoError := GetResponseInfo(mothod, URL, header, requestBytes)
	if responseInfoError != nil {
		bodyError = responseInfoError
		return bodyByte, bodyError
	}
	readAllBody, readAllError := io.ReadAll(responseInfo.Body)
	responseBodyError := responseInfo.Body.Close()
	if responseBodyError != nil {
		bodyError = responseBodyError
		return bodyByte, bodyError
	}
	if readAllError != nil {
		bodyError = readAllError
		return bodyByte, bodyError
	}
	bodyByte = readAllBody
	bodyError = nil
	return bodyByte, bodyError
}

// GetResponseBody 发起请求，返回Response body error
func GetResponseBody(mothod, URL string, header map[string][]string, requestBytes io.Reader) (*http.Response, []byte, error) {
	var result *http.Response
	var bodyByte []byte
	var bodyError error
	responseInfo, responseInfoError := GetResponseInfo(mothod, URL, header, requestBytes)
	if responseInfoError != nil {
		bodyError = responseInfoError
		return result, bodyByte, bodyError
	}
	readAllBody, readAllError := io.ReadAll(responseInfo.Body)
	responseBodyError := responseInfo.Body.Close()
	if responseBodyError != nil {
		bodyError = responseBodyError
		return result, bodyByte, bodyError
	}
	if readAllError != nil {
		bodyError = readAllError
		return result, bodyByte, bodyError
	}
	bodyByte = readAllBody
	bodyError = nil
	return responseInfo, bodyByte, bodyError
}
