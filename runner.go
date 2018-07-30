package main

import (
	"bufio"
	"log"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup
var stopChannel = make(chan bool)
var start = false

func run(startCommand string) error {

	cmd := exec.Command(startCommand)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			m := scanner.Text()
			log.Println("app |", m)
		}
	}()

	go func() {
		<-stopChannel
		log.Println("Killing Process...")
		cmd.Process.Kill()
		log.Println("Killed...")
	}()

	return nil
}

func runnerStart(context *ReuniAgentConfiguration) {
	go func() {
		log.Println("Starting program")
		err := run(context.StartCommand)
		if err != nil {
			panic(err.Error())
		}
	}()
	start = true
}

func runnerStop() {
	stopChannel <- true
}
