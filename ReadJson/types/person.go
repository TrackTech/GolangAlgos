package types

import (
	"time"
)

type Person struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	DOB  time.Time `json:"dob"`
}

type PersonList struct {
	People []Person `json:"People"`
}
