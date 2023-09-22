package main

import (
	"fmt"
)

func main3() {
	fmt.Println("You can have a inner function that repeats itself")
	var recur func(int) int
	recur = func(n int) int {
		if n == 1 {
			return 1
		}
		return n + recur(n-1)
	}

	fmt.Println(recur(5))
}
