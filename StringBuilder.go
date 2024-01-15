package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello")

	builder.WriteString("World!")

	fmt.Println(builder.String())

}
