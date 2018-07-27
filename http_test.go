package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnv() {
	os.Setenv(hostEnvVariableName, "http://127.0.0.1:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "authorizationToken")
}

func TestGetFetchUrlMethod(t *testing.T) {
	setEnv()
	initContext()
	expectedUrl := "http://127.0.0.1:8080/services/test-service/development/agent"
	assert.Equal(t, expectedUrl, getFetchVersionURL(agentConfig))
}

func TestGetConfigurationUrlMethod(t *testing.T) {
	setEnv()
	initContext()
	expectedUrl := "http://127.0.0.1:8080/services/test-service/development/2/agent"
	assert.Equal(t, expectedUrl, getFetchConfigurationURL(agentConfig, 2))
}

func TestSendRequestShouldNotReturnError(t *testing.T) {
	helper := HttpHelper{
		URL:    "http://example.com",
		Method: "GET",
	}
	resp, err := helper.SendRequest()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, resp)
}

func TestSendRequestShouldReturnError(t *testing.T) {
	helper := HttpHelper{}
	_, err := helper.SendRequest()
	assert.NotEqual(t, nil, err)
}
