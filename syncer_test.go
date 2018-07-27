package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNeedUpdate(t *testing.T) {
	assert.True(t, isNeedUpdate(1, 2))
	assert.True(t, isNeedUpdate(2, 1))
	assert.False(t, isNeedUpdate(1, 1))
}
