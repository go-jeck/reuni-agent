package main

import (
	"log"
	"time"
)

var stopLooper chan bool

func fetchVersion(config *ReuniAgentConfiguration) *Configuration {
	helper := &HttpHelper{
		URL:           getFetchVersionURL(config),
		Method:        "GET",
		Authorization: config.Authorization,
	}
	var data Configuration
	log.Println("Fetching Version...")
	err := fetchData(helper, &data)
	if err != nil {
		log.Println("Error fetching version:" + err.Error())
		return nil
	}
	log.Printf("Succesfully fetching version: %v", data.Version)
	return &data
}

func fetchConfiguration(config *ReuniAgentConfiguration, version int) *Configuration {
	helper := &HttpHelper{
		URL:           getFetchConfigurationURL(config, version),
		Method:        "GET",
		Authorization: config.Authorization,
	}
	var data Configuration
	log.Println("Fetching Configuration...")
	err := fetchData(helper, &data)
	if err != nil {
		log.Println("Error fetching configuration:" + err.Error())
		return nil
	}
	serviceConfig = &data
	log.Printf("Succesfully fetching configuration: %v", data)
	return &data
}

func isNeedUpdate(current, expected int) bool {
	return current != expected
}

func handleSync() {
	version := fetchVersion(agentConfig)
	if isNeedUpdate(serviceConfig.Version, version.Version) {
		log.Print("Configuration need to be updated")
		configuration := fetchConfiguration(agentConfig, version.Version)
		log.Println("Set configuration to environment")
		configurationSetter(configuration)
		log.Println("Configuration setted to environment")
		if start {
			stopChannel <- true
		}
		go runnerStart(agentConfig)
	} else {
		log.Print("Configuration still up to date")
	}
}

func startLooper() {
	stopLooper = make(chan bool)
	for {
		handleSync()
		select {
		case _ = <-stopLooper:
			log.Println("Stopping Looper")
			break
		default:
			time.Sleep(time.Duration(agentConfig.Interval) * time.Second)
		}
	}
}
