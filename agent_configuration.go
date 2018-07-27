package main

import (
	"errors"
	"os"
	"strconv"
)

type ReuniAgentConfiguration struct {
	Host          string `json:"host"`
	Service       string `json:"service"`
	Namespace     string `json:"namespace"`
	Authorization string `json:"authorization"`
	Interval      int    `json:"interval"`
}

const (
	hostEnvVariableName          = "REUNI_HOST"
	serviceEnvVariableName       = "REUNI_SERVICE"
	namespaceEnvVariableName     = "REUNI_NAMESPACE"
	authorizationEnvVariableName = "REUNI_AUTHORIZATION"
	intervalEnvVariableName      = "REUNI_INTERVAL"
	configErrorMessage           = "Please set up Environment variable:"
)

func isEmpty(data string) bool {
	return data == ""
}

func createConfigError(name string) error {
	return errors.New(configErrorMessage + name)
}

func initConfiguration() (*ReuniAgentConfiguration, error) {
	var config = ReuniAgentConfiguration{}
	config.Host = os.Getenv(hostEnvVariableName)
	if isEmpty(config.Host) {
		return nil, createConfigError(hostEnvVariableName)
	}
	config.Service = os.Getenv(serviceEnvVariableName)
	if isEmpty(config.Service) {
		return nil, createConfigError(serviceEnvVariableName)
	}
	config.Namespace = os.Getenv(namespaceEnvVariableName)
	if isEmpty(config.Namespace) {
		return nil, createConfigError(namespaceEnvVariableName)
	}
	config.Authorization = os.Getenv(authorizationEnvVariableName)
	if isEmpty(config.Service) {
		return nil, createConfigError(authorizationEnvVariableName)
	}
	getted := os.Getenv(intervalEnvVariableName)
	if isEmpty(getted) {
		config.Interval = 5
	} else {
		interval, err := strconv.Atoi(getted)
		if err != nil {
			return nil, createConfigError("Interval must be integer")
		}
		config.Interval = interval
	}
	return &config, nil
}
