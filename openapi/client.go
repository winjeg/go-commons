package openapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultSize = 5

var (
	// 30 seconds timeout
	httpClient = &http.Client{Timeout: time.Second * 30}
)

type Client struct {
	Key    string
	Keeper SecretKeeper
}

func NewClient(keeper SecretKeeper, key string) *Client {
	return &Client{Keeper: keeper, Key: key}
}

func DefaultClient(key, sec string) *Client {
	return &Client{
		Key: key,
		Keeper: DefaultProvider{
			AppKey:    key,
			AppSecret: sec,
		},
	}
}

func (c *Client) Get(uri string, headers ...http.Header) (string, error) {
	return c.requestWithHeader(http.MethodGet, uri, "",  headers...)
}

func (c *Client) Post(uri, body string, headers ...http.Header) (string, error) {
	return c.requestWithHeader(http.MethodPost, uri, body, headers...)
}

func (c *Client) Delete(uri, body string, headers ...http.Header) (string, error) {
	return c.requestWithHeader(http.MethodDelete, uri, body,  headers...)
}

func (c *Client) Put(uri, body string, headers ...http.Header) (string, error) {
	return c.requestWithHeader(http.MethodPut, uri, body, headers...)
}

func (c *Client) requestWithHeader(method, url, body string, headers ...http.Header) (string, error) {
	// 请求构造
	var header http.Header
	if len(headers) > 0 {
		header = headers[0]
		if len(body) >  0 {
			if header == nil {
				header = make(map[string][]string, defaultSize)
			}
			header["content-type"] = []string{"application/json"}
		}
	}
	pairs := make([]KvPair, 0, 10)
	if signHeader {
		// add all headers
		headerPairs := getPairsFromMap(header)
		pairs = append(pairs, headerPairs...)
	}
	// add all params
	timeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	paramPairs := getPairsFromUrl(url)
	pairs = append(pairs, paramPairs...)
	pairs = append(pairs, KvPair{
		Key:   timeParam,
		Value: fmt.Sprintf("%d", timeMillis),
	})
	pairs = append(pairs, KvPair{
		Key:   appKey,
		Value: c.Key,
	})
	content := buildParams(pairs)
	secret, err := c.Keeper.GetSecret(c.Key)
	if err != nil {
		return "", err
	}
	signResult := Sign(content, secret)
	if len(getPairsFromUrl(url)) == 0 {
		url += fmt.Sprintf("?sign=%s&time=%d&app_key=%s", signResult, timeMillis, c.Key)
	} else {
		url += fmt.Sprintf("&sign=%s&time=%d&app_key=%s", signResult, timeMillis, c.Key)
	}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header = header
	req.Close = true
	// 发送请求
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

func safeClose(resp *http.Response) {
	if resp != nil && !resp.Close {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("error close response body:%v", err)
		}
	}
}

func getPairsFromUrl(rawUrl string) Pairs {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil
	}
	return getPairsFromMap(u.Query())
}
