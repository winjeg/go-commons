# a tool to provide third-party apis

a common tool for providing api to third-party users  
theoretically the tool is compatible with all kinds of web framework
and `iris` and `gin` is the recommend web framework.

## for server
### before you use
if you use the default sql implementation, you should create a table first
of course, you can define your own table as long as you point out the right
way to get your actual  secret
```sql
CREATE TABLE `app` (
  `app_key` varchar(32) NOT NULL,
  `app_secret` varchar(128) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`app_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

another way for you to get the secret is you can implement the
interface here

```go
// the interface to get the secret
type SecretKeeper interface {
	GetSecret() (string, error)
}
```

### using it in your web framework
create a middle ware for some web framework
```go
// create a middle ware for iris
func OpenApiHandler(ctx iris.Context) {

    //sign header? to prevent header being modified by others
    // openapi.SignHeader(true)

	req := ctx.Request()
	// you can put the key somewhere in the header or url params
	r, err := openapi.CheckValid(req,
	// default implementation is via sql, to fetch the secrect
	    openapi.SqlSecretKeeper{
            Db:        store.GetDb(),
            TableName: "app",       // the name of table where you store all your app  keys and  secretcs
            KeyCol:    "app_key",   // the column name of the app keys
            SecretCol: "app_secret", // the column name of the app secrets
	})
	logError(err)
	if r {
	    // verfy success, continue the request
		ctx.Next()
	} else {
	    // verify fail, stop the request and return
		ctx.Text(err.Error())
		ctx.StopExecution()
		return
	}
}

```
use it on some kind of api groups
```go
// use the middle ware somewhere
// so all the apis under this group should be
// called with signed result and app key
	openApiGroup := app.Party("/open")
	openApiGroup.Use(OpenApiHandler)
	{
		openApiGroup.Get("/app", func(ctx iris.Context) {
			ctx.Text("success")
		})
	}
```

## for client
### Use client provided by the package 
```go
func GetSomeAPIResult() {
    client := openapi.DefaultClient("key....", "a9f83xa3sjh7xad")
    result, err := client.Get("/api/testdata")
    result, err := client.Post("/api/testdata", "post body")
    // ...
    // result is the data from remote api, and the remote api is enforced
    // by this very open api handler
}

```

### how to build custom clients?

1. get current time in millis and append it to the existing parameters `?time=1553759639794`
2. add `app_key`  param to your url params
3. take out all the headers and params and sort them
4. connect the sorted params to a string use `x=y&` to one string
5. sign the connected string and append the param `&sign={sign_result}` to your url parameter
6. send the request  

then you will succeed.

#### how to sign 
we only provide sha256 as the sign method of the string content
```go
// sign with sha 256
func Sign(content, key string) string {
	h := sha256.New()
	h.Write([]byte(content + key))
	return fmt.Sprintf("%x", h.Sum(nil))
}

```

#### how to sort and connect
sort order is ascending
```go
func buildParams(params Pairs) string {
	sort.Sort(params)
	var result string
	for _, v := range params {
		r := v.Key + "=" + v.Value + "&"
		result += r
	}
	return result
}
```
