package main

import (
	"os"
	"sync"
	"time"
)

var gracefulStop = make(chan os.Signal, 1)
var ticker = time.NewTicker(10 * time.Millisecond)
var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	tickerFunc()
	wg.Wait()
}
