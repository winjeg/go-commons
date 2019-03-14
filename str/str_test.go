package str

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test uuid generator to make sure it won't get duplicated keys
// in not very large range
func TestGenerateUniqueId(t *testing.T) {
	m := make(map[string]bool, 100000000)
	println(UUIDShort())
	count := 0
	for i := 0; i < 100000000; i++ {
		k := UUIDShort()
		if _, ok := m[k]; ok {
			count++
		}
		m[k] = true
	}
	if count > 0 {
		t.FailNow()
	}
}

func TestUUID(t *testing.T) {
	id1 := UUID()
	id2 := UUID()
	assert.NotEqual(t, id1, id2)
}

func TestGenerateStrings(t *testing.T) {
	fmt.Println(RandomAlphabetsLower(15))
	fmt.Println(RandomAlphabetsUpper(15))
	fmt.Println(RandomNumAlphabets(15))
	fmt.Println(RandomNumers(15))
	fmt.Println(RandomStrWithSpecialChars(15))

}

func BenchmarkKrand(b *testing.B) {
	Krand(30, KindAll)
	Krand(30, KindUpper)
	Krand(30, KindNumber)
	b.ReportAllocs()
}
