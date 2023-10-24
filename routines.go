package main

import (
	"fmt"
	"sync"
)

func processMessage(threadName string, task *sync.WaitGroup, ch <-chan int) {
	defer task.Done()
	for {
		v, ok := <-ch
		if ok {
			fmt.Println("Print value in thread (", threadName, "), v=", v)
		} else {
			fmt.Println("terminating go routing=", threadName)
			break
		}
	}
}

func main() {
	var ch chan int

	ch = make(chan int, 10)

	var arry []int = []int{1, 5, 6, 2, 3, 4, 8, 9, 7, 10}

	for _, v := range arry {
		ch <- v
	}
	close(ch)

	var wg sync.WaitGroup

	wg.Add(3)
	go processMessage("First", &wg, ch)
	go processMessage("Second", &wg, ch)
	go processMessage("Third", &wg, ch)
	wg.Wait()

}
