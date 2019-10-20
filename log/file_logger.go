// it's  not finished yet,  but all I want is a configurable and reliable logger
// 1. configurable
// 2. file logger
// 3. stdout logger
// 4. high performance
// 5. appendable

package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"sync"

	"github.com/winjeg/go-commons/str"
)

type Rolling interface {
	Duration()
}

type Appendable interface {
	Append()
}

type AppendRollingWriter interface {
	io.Writer
	Appendable
	Rolling
}

type ConfigWriter struct {
	FileName string
	Std      bool
	Lock     sync.Mutex
}

func (cw *ConfigWriter) Duration() {
}

func (cw *ConfigWriter) Append() {
}

func (cw *ConfigWriter) Write(p []byte) (n int, err error) {
	if cw.Std {
		fmt.Println(str.FromBytes(p))
	}
	if len(cw.FileName) > 0 {
		ioutil.WriteFile(cw.FileName, p, 0666)
	}
	return 0, nil
}
