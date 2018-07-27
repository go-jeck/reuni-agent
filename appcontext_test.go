package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextShouldPanic(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://127.0.0.1:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "")
	os.Setenv(authorizationEnvVariableName, "authorizationToken")
	assert.Panics(t, initContext, "initContext should panic when one environment variable not set")
}

func TestContextShouldNotPanic(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://127.0.0.1:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "authorizationToken")
	assert.NotPanics(t, initContext, "initContext should panic when one environment variable not set")
	assert.Equal(t, "development", agentConfig.Namespace)
	assert.NotEqual(t, "test", agentConfig.Service)
}
