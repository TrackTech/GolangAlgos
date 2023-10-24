/*
* 	Impleted a Ring Queue which is multiple threaded (go routines)
*	Fixed size queue buffer
*	As it is ring queue, same slot would be written once read
*	Channels used to communicate between threads
 */

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

type StoreData struct {
	Datetime  string
	Value     string
	Partition string
}

func DataProcessorRoutine(task *sync.WaitGroup, ch <-chan StoreData) {
	// The routines picks an element from ch channel
	// Currently no processing is done
	defer task.Done()
	counter := 0 // counter used to visualize how many elements proceed by this routine
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("Channel exhausted - Terminating routine.Processed = ", counter)
			break
		}
		counter++
	}
}

func DataWriterRoutine(task *sync.WaitGroup, ch chan StoreData) {
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
	var w StoreData
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
	var chanData chan StoreData
	bufferSize := 10
	args := os.Args
	if len(args) > 1 {
		size, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid buffer size. Buffer size set to default = ", bufferSize)
		} else {
			bufferSize = size
		}
	} else {
		fmt.Println("No buffer size specified. Buffer size set to default = ", bufferSize)
	}
	fmt.Println("Initializing a buffer with size = ", bufferSize)
	fmt.Println("**************************************************")
	chanData = make(chan StoreData, bufferSize)

	var waitGroup sync.WaitGroup

	waitGroup.Add(7)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataProcessorRoutine(&waitGroup, chanData)
	go DataWriterRoutine(&waitGroup, chanData)

	waitGroup.Wait()
}
