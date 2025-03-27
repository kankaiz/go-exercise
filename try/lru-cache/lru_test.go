package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	obj := Constructor(5)
	obj.Put(2, 2)
	assert.Equal(t, obj.Get(2), int(2))
}
