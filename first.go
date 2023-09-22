package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (x Vertex) Abs() float64 {

	f := (x.X * x.X) + (x.Y * x.Y)

	return math.Sqrt(f)
}

func MyGo() {
	var v Vertex
	v = Vertex{
		X: 1.0,
		Y: 2.0,
	}

	fmt.Println(v.Abs())
}
// func main() {
// 	MyGo()
// }
