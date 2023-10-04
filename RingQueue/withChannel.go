package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

type WalmartData struct {
	Datetime  string
	Value     string
	Partition string
	EOF       bool
}

func DataProcessorRoutine(task *sync.WaitGroup, ch <-chan WalmartData) {
	defer task.Done()
	for {
		element, ok := <-ch
		if !ok {
			fmt.Println("Channel exhausted-terminating routine")
			break
		}
		fmt.Println(element)
	}
}

func DataWriterRoutine(task *sync.WaitGroup, ch chan WalmartData) {
	defer task.Done()
	fileName := "data.txt"
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading file")
		close(ch)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var w WalmartData
	for {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			close(ch)
			break
		}
		json.Unmarshal(data, &w)
		ch <- w
	}
}

func main() {
	var chanData chan WalmartData
	bufferSize := 10
	args := os.Args
	if len(args) > 1 {
		bufferSize, _ = strconv.Atoi(args[1])
	}
	chanData = make(chan WalmartData, bufferSize)

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataWriterRoutine(&waitGroup, chanData)

	waitGroup.Wait()
}
