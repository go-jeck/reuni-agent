package main

import (
	"fmt"
	"runtime"
)

var count = 1

func main() {
	fmt.Printf("Thread: %v\n", runtime.GOMAXPROCS(1000))
	fmt.Printf("Thread: %v\n", runtime.GOMAXPROCS(-1))
}
