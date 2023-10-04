package main

import "fmt"

func main() {

	var mySet map[string]struct{}

	mySet = make(map[string]struct{})

	mySet["rushikesh"] = struct{}{}
	mySet["ruta"] = struct{}{}

	_, exists := mySet["rushikesh"]
	if exists {
		fmt.Println("rushikesh in set")
	}
	_, exists = mySet["asdf"]
	if !exists {
		fmt.Println("asdf NOT in set")
	}

}
