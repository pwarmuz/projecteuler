package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("started")

	Problem0002()

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}
