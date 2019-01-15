package str

import (
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
