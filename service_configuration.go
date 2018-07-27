package main

import (
	"fmt"
	"os"
)

type Configuration struct {
	Version       int                    `json:"version"`
	Configuration map[string]interface{} `json:"configuration"`
}

func configurationSetter(configurationData Configuration) {
	for k, v := range configurationData.Configuration {
		key := fmt.Sprintf("%v", k)
		value := fmt.Sprintf("%v", v)
		os.Setenv(key, value)
	}
}
