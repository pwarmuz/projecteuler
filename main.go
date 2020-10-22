package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("started")

	Problem0005()

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}
