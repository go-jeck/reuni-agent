package main

type ReuniAgentConfiguration struct {
	Host          string `json:"host"`
	Service       string `json:"service"`
	Namespace     string `json:"namespace"`
	Authorization string `json:"authorization"`
}

const (
	hostEnvVariableName          = "REUNI_HOST"
	serviceEnvVariableName       = "REUNI_SERVICE"
	namespaceEnvVariableName     = "REUNI_NAMESPACE"
	authorizationEnvVariableName = "REUNI_AUTHORIZATION"
)

func isEmpty(data string) bool {
	return data == ""
}
