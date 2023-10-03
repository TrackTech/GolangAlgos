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

type walmart struct {
	Datetime  string
	Value     string
	Partition string
	EOF       bool
}

type RingBuffer struct {
	BufferData  []walmart
	Capacity    int
	WriterIndex int
	ReaderIndex int
}

var ring RingBuffer

func Reader(task *sync.WaitGroup) {
	fmt.Println("reader called")
	//Instead of object of walmart if you stored pointer, you can check for nil
	defer task.Done()
	//counter := 0
	for {
		element := ring.BufferData[ring.ReaderIndex]
		if element.EOF == true {
			break
		}
		ring.ReaderIndex += 1
		if ring.ReaderIndex == ring.Capacity {
			ring.ReaderIndex = 0
		}
		fmt.Println(element)
	}
}

func Writer(w walmart) {
	ring.BufferData[ring.WriterIndex] = w
	ring.WriterIndex += 1
	if ring.WriterIndex == ring.Capacity {
		ring.WriterIndex = 0
	}
}

func main() {
	bufferSize := 10
	args := os.Args
	if len(args) > 1 {
		bufferSize, _ = strconv.Atoi(args[1])
	}

	ring = RingBuffer{
		Capacity:   bufferSize,
		BufferData: make([]walmart, bufferSize),
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go Reader(&waitGroup)

	fileName := "data.txt"

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("errr")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			wEof := walmart{
				EOF: true,
			}
			Writer(wEof)
			fmt.Println("done reader")
			break
		}
		//fmt.Println(data)
		// counter += 1
		// if counter == 10 {
		// 	break
		// }
		w := walmart{}
		json.Unmarshal(data, &w)
		//fmt.Println(w)
		Writer(w)
	}
	waitGroup.Wait()
}
