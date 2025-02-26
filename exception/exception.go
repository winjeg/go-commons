package exception

import "github.com/winjeg/go-commons/log"

// Handle catch the exception and use the customized function handler to
// process the exception
func Handle(handler func(e any)) {
	if err := recover(); err != nil {
		handler(err)
	}
	return
}

// Catch to catch exception
// and log it
func Catch() {
	if err := recover(); err != nil {
		log.GetLogger(nil).Errorf("Catch exception occured: %+v", err)
	}
	return
}
