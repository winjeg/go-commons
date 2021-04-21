# httpClient
a simple client to make remote http request  and return string response body

currently supported method:
1. Get
2. Put   (content type: json)
3. Post (content type: json)
4. Delete (content type: json)


## example code

```go
package any

import (
    "github.com/stretchr/testify/assert"
	"github.com/winjeg/go-commons/httpclient"
)

func TestGet(t *testing.T) {
	d, err := httpclient.Get("https://www.baidu.com")
	assert.Equal(*t, err == nil, true)
	assert.Equal(*t, len(d) > 0, true)
	
	_, err = httpclient.Post("https://www.baidu.com", `{"abc":"def"}`)
	assert.NotNil(*t, err)
}
```
