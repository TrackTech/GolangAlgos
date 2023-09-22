package main

import (
	"fmt"
	"sync"
)

var ch chan int

func processMessage(task *sync.WaitGroup) {
	defer task.Done()
	v := <-ch

	fmt.Println("Print value=", v)
}

func main10() {

	ch = make(chan int, 10)

	var arry []int = []int{1, 5, 6, 2, 8, 9, 10}

	for _, v := range arry {
		ch <- v
	}

	var wg sync.WaitGroup

	index := 0
	for index < 7 {
		wg.Add(1)
		go processMessage(&wg)
	}
	wg.Wait()

}
