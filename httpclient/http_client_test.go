package httpclient

import (
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	d, err := Get("https://www.baidu.com")
	assert.Equal(t, err == nil, true)
	assert.Equal(t, len(d) > 0, true)
}

func TestGetWithParams(t *testing.T) {
	m := make(map[string]string, 2)
	m["tn"] = "baidu"
	m["wd"] = "s"
	d, err := GetWithParams("https://www.baidu.com/s", m)
	assert.Equal(t, err == nil, true)
	assert.Equal(t, len(d) > 0, true)
	w, err := GetWithParams("https://www.baidu.com/sasdasdas", nil)
	assert.Equal(t, strings.Index(w, "页面不存在") > -1, true)
}
