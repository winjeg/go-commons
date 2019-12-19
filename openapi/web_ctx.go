package openapi

import (
	"errors"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultExpireTime = 60000
	signParam         = "sign"
	timeParam         = "time"
	appKey            = "app_key"
)

var signHeader = false

// SignHeader whether to sign http request header or not
func SignHeader(s bool) {
	var lock sync.Mutex
	lock.Lock()
	signHeader = s
	lock.Unlock()
}

// CheckValid to check if the request is valid  from the signing key
func CheckValid(req *http.Request, keeper SecretKeeper) (bool, error) {
	if req == nil {
		return false, errors.New("illegal request")
	}
	// time in millis
	timeStr := getParamFromRequest(req, timeParam)
	signResult := getParamFromRequest(req, signParam)

	rt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return false, errors.New("error parameter")
	}
	now := time.Now().UnixNano() / int64(time.Millisecond)
	duration := math.Abs(float64(rt - now))
	if duration > defaultExpireTime {
		return false, errors.New("error timestamp")
	}

	pairs := getPairs(req)
	content := buildParams(pairs)
	secret, err := keeper.GetSecret(getParamFromRequest(req, appKey))
	if err != nil {
		return false, err
	}
	result := verify(signResult, content, secret)
	if result {
		return result, nil
	}
	return result, errors.New("error verifying")
}

func Do(req *http.Request, keeper SecretKeeper, successFun, failFunc func()) {
	if r, err := CheckValid(req, keeper); r {
		successFun()
	} else {
		log.Printf("%v\n", err)
		failFunc()
	}
}

func getParamFromRequest(req *http.Request, param string) string {
	if req == nil {
		return ""
	}
	return req.URL.Query().Get(param)
}

func getPairs(req *http.Request) Pairs {
	pairs := make([]KvPair, 0, 10)
	if signHeader {
		// add all headers
		headers := req.Header
		headerPairs := getPairsFromMap(headers)
		pairs = append(pairs, headerPairs...)
	}

	// add all params
	paramsMap := req.URL.Query()
	paramPairs := getPairsFromMap(paramsMap)
	pairs = append(pairs, paramPairs...)
	return pairs
}

// get params and headers except the param sign
func getPairsFromMap(m map[string][]string) Pairs {
	pairs := make([]KvPair, 0, 10)
	for k, v := range m {
		if len(k) < 1 {
			continue
		}
		var val string
		for _, e := range v {
			val += e
		}
		if strings.EqualFold(k, signParam) {
			continue
		}
		p := KvPair{
			Key:   k,
			Value: val,
		}
		pairs = append(pairs, p)
	}
	return pairs
}
