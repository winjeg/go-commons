package exception

import (
	"fmt"
	"log"
	"testing"
)

func TestHandle(t *testing.T) {
	defer Handle(func(e any) {
		log.Printf("error occured: %+v\n", e)
	})
	x := 0
	fmt.Println(3 / x)
}

func TestCatch(t *testing.T) {
	defer Catch()
	x := 0
	fmt.Println(3 / x)
}
