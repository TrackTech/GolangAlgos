package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func (this Person) Object_Receiver() {
	this.Name = "Name change"
	fmt.Println("The object passed does not change")
}

func (this *Person) Pointer_Receiver() {
	this.Name = "Olga"
	fmt.Printf("Pointer changes the value of the object passed")
}

// func main() {
// 	var p Person
// 	p = Person{
// 		Name:   "Rushikesh",
// 		Age:    42,
// 		Salary: 0.0,
// 	}
// 	p.Object_Receiver()
// 	fmt.Println(p)
// 	p.Pointer_Receiver()
// 	fmt.Println(p)
// }
