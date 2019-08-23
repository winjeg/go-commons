package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextID(t *testing.T) {
	assert.NotEqual(t, NextID(), NextID())
	assert.True(t, NextID() < NextID())
}

func BenchmarkNextID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextID()
	}
}

func TestUUID(t *testing.T) {
	assert.NotEqual(t, UUID(), UUID())
	assert.NotEqual(t, UUIDShort(), UUIDShort())
}
