package main

import (
	"fmt"
	"reflect"
)

func main2() {

	s := "rushikesh"

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Println(reflect.TypeOf(s[i]))
		break
	}

	for _, ch := range s {
		fmt.Println(ch)
		fmt.Println(reflect.TypeOf(ch))
		break
	}
}
