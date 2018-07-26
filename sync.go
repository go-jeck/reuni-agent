package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type configuration struct {
	Version       int               `json:"version"`
	Configuration map[string]string `json:"configuration"`
}

func tickerFunc() {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		fetchVersion()
		for _ = range ticker.C {
			fetchVersion()
		}
	}()
}

func fetchVersion() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/services/go-pay-payment/dev-box2/2/agent", nil)
	req.Header.Set("Authorization", "op98w6zTca9DPHx2pAH0kEBxXBpjcVMF7kAEp+2xenwkxSxAeQhU541US8hNTmb2")

	resp, _ := client.Do(req)
	fmt.Println(resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
}
