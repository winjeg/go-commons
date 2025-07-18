package hash

import (
	"fmt"
	"testing"
)

func TestXXHash(t *testing.T) {
	fmt.Println(XXHashSum64([]byte("Hello1")))
	fmt.Println(XXHashSum64([]byte("Hello2")))
	fmt.Println(XXHashSum64String("Hello1"))
	fmt.Println(XXHashSum64String("Hello2"))
}
