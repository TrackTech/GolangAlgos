package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(task *sync.WaitGroup, stopChan <-chan struct{}) {
	defer task.Done()
	loop := true

	for loop {
		select {
		case <-stopChan:
			fmt.Println("Stop Channel termination message received")
			loop = false
		default:
			now := time.Now()
			fmt.Println(now)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	var waitGroup sync.WaitGroup
	var stopChan chan struct{}

	stopChan = make(chan struct{})

	waitGroup.Add(1)

	go reader(&waitGroup, stopChan)

	for {
		currSecond := time.Now().Second()
		if currSecond == 59 {
			close(stopChan)
			break
		}
	}

	waitGroup.Wait()

}
