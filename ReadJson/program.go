package main

import (
	"ReadJson/reader"
	"ReadJson/types"
	"ReadJson/writer"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "test.json"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	reader.ReadFile(fileName)

	family := []types.Person{}

	family = append(family, types.Person{
		Name: "Rushikesh",
		Age:  42,
		DOB:  time.Date(1981, time.March, 25, 0, 0, 0, 0, time.UTC),
	})
	family = append(family, types.Person{
		Name: "Olga",
		Age:  40,
		DOB:  time.Date(1983, time.February, 12, 0, 0, 0, 0, time.UTC),
	})

	err := writer.WriteJson("test2.json", family)

	if err != nil {
		fmt.Println("Error happened")
		fmt.Println(err)
	}

}
