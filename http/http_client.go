package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/winjeg/go-commons/log"
)

var (
	// 3 seconds timeout
	httpClient      = &http.Client{Timeout: time.Second * 30}
	jsonContentType = []string{"application/json"}
)

const (
	contentTypeHeaderName = "content-type"
)

// export GetWithParams
// Get request with url params
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

func GetWithHeader(url string, header http.Header) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header = header
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		return string(data), errors.New(fmt.Sprintf("Error with not correct status code %s", resp.Status))
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

// Put request by url
// content type is application/json
func Put(url, content string) (string, error) {
	logger := log.GetLogger(nil)
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(content))
	req.Header = http.Header{
		contentTypeHeaderName: jsonContentType,
	}
	if err != nil {
		return "", err
	}
	req.Close = true
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

// export Delete
// delete request by url
// content type is application/json
func Delete(url, content string) (string, error) {
	logger := log.GetLogger(nil)
	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(content))
	req.Header = http.Header{
		contentTypeHeaderName: jsonContentType,
	}
	if err != nil {
		return "", err
	}
	req.Close = true
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

// export Delete
// delete request by url
// content type is application/json
func DoRequest(method, url, content string, header http.Header) (string, error) {
	logger := log.GetLogger(nil)
	req, err := http.NewRequest(method, url, strings.NewReader(content))
	req.Header = header
	if err != nil {
		return "", err
	}
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Errorf("%v\n", err)
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
		if err != nil {
			logger.Errorf("error close response body:%v\n", err)
		}
	}
}
