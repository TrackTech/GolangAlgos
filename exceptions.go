package main

import (
	"fmt"
	"strconv"
)

type CustomError struct {
	Message string
	Code    int
}

func (c CustomError) Error() {
	fmt.Println(c.Message + " " + strconv.Itoa(c.Code))
}

func divide_by_zero(x, y int) (int, CustomError) {
	if y == 0 {
		return -1, CustomError{
			Message: "Nice Try",
			Code:    -1,
		}
	} else {
		return x / y, *&CustomError{}
	}
}

func main7() {
	x := 10
	y := 0

	divide_by_zero(x, y)
}
