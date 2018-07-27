package main

var agentConfig *ReuniAgentConfiguration

func initContext() {
	var err error
	agentConfig, err = initConfiguration()
	if err != nil {
		panic(err)
	}
}
