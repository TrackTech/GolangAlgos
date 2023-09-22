package main

import (
	"fmt"
	"strconv"
)

func main5() {
	n := 15
	bin := strconv.FormatInt(int64(n), 2)
	fmt.Println("15 bin =", bin)
	n = -15
	bin = strconv.FormatInt(int64(n), 2)
	fmt.Println("-15 bing =", bin)
}
