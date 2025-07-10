package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	c := Filter([]string{"a", "b", "c"}, func(arg string) bool {
		return arg == "b"
	})
	assert.True(t, len(c) == 1)
	assert.True(t, c[0] == "b")
	assert.NotNil(t, Filter([]string{}, func(string) bool { return false }))
	assert.Nil(t, Filter(nil, func(string) bool { return false }))
}

func TestFilterMap(t *testing.T) {
	m := FilterMap(map[int]string{1: "a", 2: "b"}, func(k int, v string) bool {
		return k > 1 && len(v) == 1
	})
	assert.True(t, len(m) == 1)
	assert.True(t, m[2] == "b")
	assert.Nil(t, FilterMap(nil, func(k string, v int) bool { return true }))
	assert.NotNil(t, FilterMap(map[int]int{}, func(k int, v int) bool { return false }))
}
