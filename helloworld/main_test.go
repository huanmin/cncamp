package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrease(t *testing.T) {
	t.Log("Starting testing")
	result := add(1, 2)
	assert.Equal(t, result, 3)
}
