package main

import (
	"errors"
	"os"
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
	err := createConfigError("test")
	assert.Equal(t, err, errors.New(configErrorMessage+"test"))
}

func TestInitConfigShouldNotReturnError(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://localhost:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "123456")
	config, err := initConfiguration()
	assert.NotEqual(t, nil, config, "configuration should not be nil")
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, "123456", config.Authorization)
	assert.Equal(t, "development", config.Namespace)
	assert.Equal(t, "http://localhost:8080", config.Host)
	assert.Equal(t, "test-service", config.Service)
}

func TestInitConfigShouldReturnError(t *testing.T) {
	os.Setenv(hostEnvVariableName, "")
	os.Setenv(serviceEnvVariableName, "")
	os.Setenv(namespaceEnvVariableName, "")
	os.Setenv(authorizationEnvVariableName, "")
	_, err := initConfiguration()
	assert.NotEqual(t, nil, err)
}

func TestInitConfigShouldReturnErrorWhenIntervalNotInteger(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://localhost:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "123456")
	os.Setenv(intervalEnvVariableName, "asd")

	_, err := initConfiguration()
	assert.Error(t, err)

}

func TestInitConfigShouldNotReturnErrorWhenIntervalNotSet(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://localhost:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "123456")
	os.Setenv(intervalEnvVariableName, "")

	config, err := initConfiguration()
	assert.NoError(t, err)
	assert.Equal(t, config.Interval, 5)
}

func TestInitConfigShouldNotReturnErrorWhenIntervalSetted(t *testing.T) {
	os.Setenv(hostEnvVariableName, "http://localhost:8080")
	os.Setenv(serviceEnvVariableName, "test-service")
	os.Setenv(namespaceEnvVariableName, "development")
	os.Setenv(authorizationEnvVariableName, "123456")
	os.Setenv(intervalEnvVariableName, "60")

	config, err := initConfiguration()
	assert.NoError(t, err)
	assert.Equal(t, config.Interval, 60)

}
