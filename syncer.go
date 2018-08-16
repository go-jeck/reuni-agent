package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis"
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

func handleSync(pubSubClient *redis.PubSub) {
	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubSubClient.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubSubClient.Channel()

	getLatestConfig(fetchVersion(agentConfig).Version)

	// go func() {
	for {
		msg, ok := <-ch

		if !ok {
			break
		}
		fmt.Println(msg.Channel, msg.Payload)
		log.Print("Configuration need to be updated")
		version, err := strconv.Atoi(msg.Payload)
		if err != nil {
			log.Println("ERROR PARSING VERSION")
		} else {
			getLatestConfig(version)
		}
	}
	// }()

}

func getLatestConfig(version int) {
	// version := fetchVersion(agentConfig)
	// if isNeedUpdate(serviceConfig.Version, version.Version) {
	log.Print("Configuration need to be updated")
	configuration := fetchConfiguration(agentConfig, version)
	log.Println("Set configuration to environment")
	configurationSetter(configuration)
	log.Println("Configuration setted to environment")
	if start {
		stopChannel <- true
	}
	go runnerStart(agentConfig)
	// } else {
	// 	log.Print("Configuration still up to date")
	// }
}

// func startLooper(pubSubClient *redis.PubSub) {
// 	stopLooper = make(chan bool)
// 	for {
// 		handleSync(pubSubClient)
// 		select {
// 		case _ = <-stopLooper:
// 			log.Println("Stopping Looper")
// 			break
// 		default:
// 			time.Sleep(time.Duration(agentConfig.Interval) * time.Second)
// 		}
// 	}
// }
