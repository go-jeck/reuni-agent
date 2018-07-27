package main

import (
	"fmt"
	"net/http"
)

type HttpCaller interface {
	SendRequest() *http.Response
}

func getFetchVersionURL(config *ReuniAgentConfiguration) string {
	return fmt.Sprintf("%v/services/%v/%v/agent", config.Host, config.Service, config.Namespace)
}
