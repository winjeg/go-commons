package openapi

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckValid(t *testing.T) {
	SignHeader(false)
	time := time.Now().UnixNano() / 1e6
	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("http://localhost:8080?time=%d&app_key=thekey&id=123", time),
		nil)
	pairs := getPairs(req)
	sec, err := sqlKeeper.GetSecret("thekey")
	assert.Nil(t, err)
	signResult := Sign(buildParams(pairs), sec)
	req.URL.RawQuery += "&sign=" + signResult
	assert.Nil(t, err)
	_, err = CheckValid(req, sqlKeeper)
	assert.Nil(t, err)
	req.URL.RawQuery += "&abc=1"
	_, err = CheckValid(req, sqlKeeper)
	assert.NotNil(t, err)
	SignHeader(false)
	_, err = CheckValid(req, sqlKeeper)
	assert.NotNil(t, err)
}
