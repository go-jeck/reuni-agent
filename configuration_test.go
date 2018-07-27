package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyShouldReturnTrueWithEmptyString(t *testing.T) {
	assert.Equal(t, isEmpty(""), true)
}

func TestIsEmptyShouldReturnFalseWithNotEmptyString(t *testing.T) {
	assert.Equal(t, isEmpty("test"), false)
}
