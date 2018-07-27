package main

import (
	"fmt"
	"net/http"
)

type HttpCaller interface {
	SendRequest() *http.Response
}

type HttpHelper struct {
	URL           string
	Method        string
	Authorization string
}

func getFetchVersionURL(config *ReuniAgentConfiguration) string {
	return fmt.Sprintf("%v/services/%v/%v/agent", config.Host, config.Service, config.Namespace)
}

func getFetchConfigurationURL(config *ReuniAgentConfiguration, version int) string {
	return fmt.Sprintf("%v/services/%v/%v/%v/agent", config.Host, config.Service, config.Namespace, version)
}

func (h *HttpHelper) SendRequest() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(h.Method, h.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", h.Authorization)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
