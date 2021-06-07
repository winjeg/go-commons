package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	id1 := UUID()
	id2 := UUID()
	assert.NotEqual(t, id1, id2)
}

func TestGenerateStrings(t *testing.T) {
	fmt.Println(RandomAlphabetsLower(15))
	fmt.Println(RandomAlphabetsUpper(15))
	fmt.Println(RandomNumAlphabets(15))
	fmt.Println(RandomNumbers(15))
	fmt.Println(RandomStrWithSpecialChars(15))
	str := UUIDShort()
	assert.NotNil(t, str)
	assert.True(t, !IsBlank(str))
	assert.True(t, !IsEmpty(str))
	r := ReplaceAll(str, "a", "A")
	assert.NotNil(t, r)
}

func BenchmarkKrand(b *testing.B) {
	RandStr(30, KindAll)
	RandStr(30, KindUpper)
	RandStr(30, KindNumber)
	b.ReportAllocs()
}

// test uuid generator to make sure it won't get duplicated keys
// in not very large range
func BenchmarkGenerateUniqueId(b *testing.B) {
	m := make(map[string]bool, 1000000)
	println(UUIDShort())
	count := 0
	for i := 0; i < 100000000; i++ {
		k := UUIDShort()
		if _, ok := m[k]; ok {
			count++
		}
		m[k] = true
	}
	b.ReportAllocs()
}

func TestTrim(t *testing.T) {
	a := ",abc,efg.,"
	b := TrimComma(a)
	assert.Equal(t, b, "abc,efg.")
	c := TrimDot(b)
	assert.Equal(t, "abc,efg", c)
	d := Trim(a, ".,")
	assert.Equal(t, d, ",abc,efg")
}

func BenchmarkJoin(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Join([]string{"a", "b"}, "c")
	}
}

func TestConvert(t *testing.T) {
	a := "123中国人"
	bts := ToBytes(a)
	b := FromBytes(bts)
	assert.Equal(t, a, b)
}

func BenchmarkConvert(b *testing.B) {
	a := "123中国人"
	for i := 0; i < b.N; i++ {
		bts := ToBytes(a)
		FromBytes(bts)
	}
	b.ReportAllocs()
}

func TestJoin(t *testing.T) {
	joinStr := JoinInt([]int{1, 2}, ",")
	assert.Equal(t, joinStr, "1,2")
	arr, err := SplitInt(joinStr, ",")
	assert.Nil(t, err)
	assert.Len(t, arr, 2)
	arr2, err2 := SplitInt64(joinStr, ",")
	assert.Nil(t, err2)
	assert.Len(t, arr2, 2)
	result := JoinIfNotEmpty(",", "a", "", "c")
	assert.Equal(t, result, "a,c")
}

func TestIsAllBlank(t *testing.T) {
	assert.True(t, IsAllBlank(""))
	assert.True(t, IsAllBlank())
	assert.True(t, IsAllBlank("", "     "))
	assert.False(t, IsAllBlank("", "     ", "13"))
	assert.True(t, IsNotBlank(" _"))
	assert.False(t, IsNoneBlank())
	assert.False(t, IsNoneBlank(" "))
	assert.False(t, IsNoneBlank(" ", ""))
	assert.True(t, IsNoneBlank(" a", "1"))
	assert.True(t, HasAnyBlank())
	assert.True(t, HasAnyBlank(""))
	assert.True(t, HasAnyBlank("", "a"))
	assert.False(t, HasAnyBlank("c", "a"))
	assert.True(t, NotAllBlank("c", "a"))
	assert.True(t, NotAllBlank("", "a"))
	assert.False(t, NotAllBlank("", " "))
	assert.False(t, NotAllBlank())
}
