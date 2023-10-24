package main

import "fmt"

func main() {

	var mySet map[string]struct{}

	mySet = make(map[string]struct{})

	mySet["Golang"] = struct{}{}
	mySet["Python"] = struct{}{}

	_, exists := mySet["Golang"]
	if exists {
		fmt.Println("Golang in set")
	}
	_, exists = mySet["asdf"]
	if !exists {
		fmt.Println("asdf NOT in set")
	}

}
