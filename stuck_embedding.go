package main

import "fmt"

type address struct {
	city  string
	state string
}

type university struct {
	name       string
	chanceller string
}

type student struct {
	name               string
	student_address    address
	student_university *university
}

func (this student) String() string {

	s := fmt.Sprintf("Name=%s, city=%s,university=%s", this.name, this.student_address.city, this.student_university.name)
	return s

}

func main() {

	var class []student
	class = []student{}
	standord := university{
		name:       "Stanfor",
		chanceller: "Me",
	}

	class = append(class, student{
		name: "Rushikesh",
		student_address: address{
			city:  "san jose",
			state: "ca",
		},
		student_university: &standord,
	})
	for _, s := range class {
		fmt.Println(s)
	}
}
