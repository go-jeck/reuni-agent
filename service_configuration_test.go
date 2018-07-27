package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationEnvironmentSetter(t *testing.T) {
	configuration := make(map[string]interface{})
	configuration["DB_HOST"] = "localhost"
	configuration["DB_PASS"] = "test123"
	configData := Configuration{
		Version:       10,
		Configuration: configuration,
	}
	configurationSetter(&configData)
	assert.Equal(t, "localhost", os.Getenv("DB_HOST"))
	assert.Equal(t, "test123", os.Getenv("DB_PASS"))
}
