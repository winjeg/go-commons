package main

import (
	"fmt"
	"github.com/winjeg/go-commons/httpclient"
)

func main() {
	_, err := httpclient.Get("https://www.bing.com")
	if err != nil {
		fmt.Println(err.Error())
	}
}
