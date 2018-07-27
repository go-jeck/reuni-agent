package main

type Configuration struct {
	Version       int                    `json:"version"`
	Configuration map[string]interface{} `json:"configuration"`
}
