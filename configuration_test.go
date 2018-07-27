package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyShouldReturnTrueWithEmptyString(t *testing.T) {
	assert.Equal(t, isEmpty(""), true)
}

func TestIsEmptyShouldReturnFalseWithNotEmptyString(t *testing.T) {
	assert.Equal(t, isEmpty("test"), false)
}

func TestCreateConfigError(t *testing.T) {
	err := createConfigError("REUNI_HOST")
	assert.Equal(t, err, errors.New(configErrorMessage+"REUNI_HOST"))
}
