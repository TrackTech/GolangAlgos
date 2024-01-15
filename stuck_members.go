package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type People struct {
	id     int
	name   string
	salary float64
	D      time.Time
}

func mains() {
	fmt.Printf("Hello world")

	file, err := os.Open(".txt")
	if err != nil {
		fmt.Println("bad file")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("==done=-")
			break
		}
		words := strings.Split(string(line), " ")

		id, _ := strconv.Atoi(words[0])
		layout := "2006-01-02"
		dob, _ := time.Parse(layout, words[3])
		fmt.Println("converted date", words[3], "--", dob)
		fmt.Println(reflect.TypeOf(dob))
		p := People{
			id:   id,
			name: words[1],
			D:    dob,
		}
		fmt.Println("p is ", p)
	}
}
