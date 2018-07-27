package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

type MockHTTPCaller struct {
	Response *http.Response
}

func (h *MockHTTPCaller) SendRequest() (*http.Response, error) {
	return h.Response, nil
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

func TestFetchDataShouldReturnErrorWithResponse404(t *testing.T) {
	resp := &http.Response{
		StatusCode: 404,
		Status:     "404 Not Found",
	}
	caller := &MockHTTPCaller{
		Response: resp,
	}
	var data Configuration
	err := fetchData(caller, &data)
	assert.EqualError(t, err, "HTTP Error: 404 Not Found")
	assert.Equal(t, 0, data.Version)
}

func TestFetchDataShouldReturnErrorWithResponse500(t *testing.T) {
	resp := &http.Response{
		StatusCode: 500,
		Status:     "500 Internal Server Error",
	}
	caller := &MockHTTPCaller{
		Response: resp,
	}
	var data Configuration
	err := fetchData(caller, &data)
	assert.EqualError(t, err, "HTTP Error: 500 Internal Server Error")
	assert.Equal(t, 0, data.Version)
}

func TestFetchDataShouldResponseErrorWhenBodyIsNotValidJSON(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(bytes.NewBufferString("Hello World")),
	}
	caller := &MockHTTPCaller{
		Response: resp,
	}
	var data Configuration
	err := fetchData(caller, &data)
	assert.Error(t, err)
}

func TestFetchDataShouldNotReturnErrorWhenResponseBodyIsValidJSON(t *testing.T) {
	configuration := make(map[string]interface{})
	configuration["DB_HOST"] = "localhost"
	configuration["DB_PASS"] = "test123"
	responseData := Configuration{
		Version:       10,
		Configuration: configuration,
	}
	responseJSON, err := json.Marshal(responseData)
	resp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(bytes.NewBufferString(string(responseJSON))),
	}
	caller := &MockHTTPCaller{
		Response: resp,
	}
	var data Configuration
	err = fetchData(caller, &data)
	assert.Equal(t, nil, err)
	assert.Equal(t, data.Version, 10)
	assert.Equal(t, data.Configuration, configuration)
}
