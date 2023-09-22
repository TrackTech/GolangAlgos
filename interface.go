package main

import (
	"fmt"
)

type Vehicle interface {
	get_type() string
	model() string
}

func get_vehicle_info(v Vehicle) {
	ty := v.get_type()
	mo := v.model()
	fmt.Println("Vehicle type is =", ty, ", model is=", mo)
}

type Honda struct {
	type_name  string
	model_name string
}

type Toyota struct {
	t_name string
	m_name string
}

func (this Honda) get_type() string {
	return this.type_name
}
func (this Honda) model() string {
	return this.model_name
}

func (this Toyota) get_type() string {
	return this.t_name
}
func (this Toyota) model() string {
	return this.m_name
}

// func main() {
// 	H := Honda{
// 		type_name:  "Honda",
// 		model_name: "Civic",
// 	}
// 	var T Toyota
// 	T = Toyota{
// 		t_name: "Toyota",
// 		m_name: "Prius",
// 	}
// 	get_vehicle_info(H)
// 	get_vehicle_info(T)
// }
