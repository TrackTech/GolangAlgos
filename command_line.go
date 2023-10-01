package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	args := os.Args
	fileName := "test.json"
	if len(args) > 1 {
		fileName = args[1]
	}
	fmt.Println(fileName)
}
