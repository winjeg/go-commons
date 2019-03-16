# props
property parser for go
## 

##  how to install
```go
go get github.com/winjeg/props
```
## samples
```go

package main
import "github.com/winjeg/props"

const propText = `
username=tester
`
func main() {
    // load from file
    props.LoadFile("test.props")
    // from text
    p := props.FromString(propText)
    // back to props string
    println(p.String())
}
```

## notice
please don't fork this project, but contribute to this project.


