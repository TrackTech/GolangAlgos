package main

import (
	"errors"
	"fmt"
	"strconv"
)

var MultipeByZeroErr = errors.New("Mutiple by zero")

//errors.New("mutiple by zero")

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

func multiply_by_zero(x, y int) (int, error) {
	if x == 0 || y == 0 {
		return -1, MultipeByZeroErr
	}
	return x * y, nil
}

func main7() {
	x := 10
	y := 0

	divide_by_zero(x, y)
}
