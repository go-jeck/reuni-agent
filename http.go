package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpCaller interface {
	SendRequest() (*http.Response, error)
}

type HttpHelper struct {
	URL           string
	Method        string
	Authorization string
}

func getFetchVersionURL(config *ReuniAgentConfiguration) string {
	return fmt.Sprintf("%v/%v/%v/%v/agent", config.Host, config.Organization, config.Service, config.Namespace)
}

func getFetchConfigurationURL(config *ReuniAgentConfiguration, version int) string {
	return fmt.Sprintf("%v/%v/%v/%v/%v/agent", config.Host, config.Organization, config.Service, config.Namespace, version)
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

func fetchData(caller HttpCaller, data interface{}) error {
	resp, err := caller.SendRequest()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP Error: " + resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
