package httpclient

import (
	"errors"
	"github.com/winjeg/go-commons/log"
	"strconv"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 5秒的超时时间
var (
	httpClient = &http.Client{Timeout: 5e9}
)

// 对其进行简单封装， 可以传参数列表的map
func GetWithParams(uri string, paramMap map[string]string) (string, error) {
	url := uri + "?"
	for k, v := range paramMap {
		url += fmt.Sprintf("%s=%s&", k, v)
	}
	lastLetterIdx := len(url) - 1
	qIdx := strings.LastIndex(url, "?")
	andIdx := strings.LastIndex(url, "&")

	if qIdx == lastLetterIdx || andIdx == lastLetterIdx {
		url = url[:lastLetterIdx]
	}
	return Get(url)
}

// Get request by url
func Get(url string) (string, error) {
	resp, err := httpClient.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error with status code:" + strconv.Itoa(resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer safeClose(resp)
	return string(data), err
}

// Post request by url
// content type is application/json
func Post(url, content string) (string, error) {
	logger := log.GetLogger(nil)
	resp, err := httpClient.Post(url, "application/json", strings.NewReader(content))
	if err != nil {
		logger.Errorf("Error post to %s, error : %v", url, err)
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error with status code:" + strconv.Itoa(resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer safeClose(resp)
	return string(data), err
}

// Post request by url
// content type is application/json
func Put(url, content string) (string, error) {
	logger := log.GetLogger(nil)
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(content))
	req.Header = http.Header{
		"content-type": {"application/json"},
	}
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Errorf("%v", err)
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error with status code:" + strconv.Itoa(resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer safeClose(resp)
	return string(data), err
}

// Post request by url
// content type is application/json
func Delete(url, content string) (string, error) {
	logger := log.GetLogger(nil)
	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(content))
	req.Header = http.Header{
		"content-type": {"application/json"},
	}
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Errorf("%v", err)
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error with status code:" + strconv.Itoa(resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer safeClose(resp)
	return string(data), err
}

func safeClose(resp *http.Response) {
	logger := log.GetLogger(nil)
	if resp != nil && !resp.Close {
		err := resp.Body.Close()
		logger.Errorf("error close response body:%v", err)
	}
}
