package main

var agentConfig *ReuniAgentConfiguration
var serviceConfig *Configuration

func initContext() {
	var err error
	agentConfig, err = initConfiguration()
	if err != nil {
		panic(err)
	}
}
