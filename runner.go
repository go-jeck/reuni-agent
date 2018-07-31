package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup
var stopChannel = make(chan bool)
var start = false

func run(startCommand string, arg string) error {

	cmd := exec.Command(startCommand, arg)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
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

	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			m := errScanner.Text()
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
		err := run(context.StartCommand, os.Getenv("REUNI_ARGS"))
		if err != nil {
			panic(err)
		}
	}()
	start = true
}

func runnerStop() {
	stopChannel <- true
}
