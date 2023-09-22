package main

import (
	"fmt"
)

func outer() {
	var ret int
	ret = 10

	var inner func()

	inner = func() {
		ret := 100
		fmt.Println("I can access outer variables , ret=", ret)
		ret += 10
	}
	inner()
	fmt.Println("Inner functions can modify ret, ret=", ret)
}

func main6() {
	outer()
}
