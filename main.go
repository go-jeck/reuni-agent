package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!!")
	fmt.Println(os.Getenv("REUNI_HOST"))
	fmt.Println(os.Getenv("REUNI_SERVICE"))
	fmt.Println(os.Getenv("REUNI_NAMESPACE"))
	fmt.Println(os.Getenv("REUNI_AUTHORIZATION"))
}
