package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testKey = "key1"
	testValue = "val2"
	testValue2 = "Val2"
	testExp = 5
)

func TestGet(t *testing.T) {
	assert.Nil(t, Get(testKey))
	Set(testKey, testValue, testExp)
	assert.Equal(t, Get(testKey), testValue)
	time.Sleep(time.Millisecond * testExp)
	assert.Nil(t, Get(testKey))
	Set(testKey, testValue2, testExp)
	assert.Equal(t, Get(testKey), testValue2)
}