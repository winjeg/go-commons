package openapi

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
	"time"
)

const (
	signData   = "Hello, ]workld"
	signKey    = "12jhdjx-1233xahsd-313ksjkx"
	signResult = "863fc99a95f752c83ca63e716fa9e83cfa35fa06a158802c60ac30a5bb734f9c"
)

func TestSign(t *testing.T) {
	result := Sign(signData, signKey)
	assert.Equal(t, result, signResult)
}

func TestVerify(t *testing.T) {
	r := verify(signResult, signData, signKey)
	assert.Equal(t, r, true)
}

func TestBuildParams(t *testing.T) {
	time := time.Now().UnixNano() / 1e6
	var pairs Pairs
	pairs = append(pairs, KvPair{
		Key:   "time",
		Value: fmt.Sprintf("%d", time),
	})
	pairs = append(pairs, KvPair{
		Key:   "app_key",
		Value: "iepheechahNg9voh",
	})
	content := buildParams(pairs)
	url := "?time=" + fmt.Sprintf("%d", time) + "&app_key=iepheechahNg9voh&sign=" + Sign(content, signKey)
	fmt.Println(url)
}
