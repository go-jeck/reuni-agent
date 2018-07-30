package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var stopSignal chan os.Signal

func initGracefulShutdown() {
	log.Println("Set up shut down handler")
	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopSignal
		log.Println("Gracefully shutting down agent")

		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
}
