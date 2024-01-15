package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func processData(name string, task *sync.WaitGroup, ch <-chan int) {
	defer task.Done()

	for {
		d, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Tread - ", name, ", processed-", d)
	}
}

func processDataPanic(name string, task *sync.WaitGroup, ch <-chan int) {
	// defer task.Done()
	defer func() {
		task.Done()
		for i := 0; i < 4; i++ {
			fmt.Println("*****DEFER******")
		}
		fmt.Println("I am done executing with defer-", name)
	}()
	var builder strings.Builder
	counter := 0
	for {
		d, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Tread - ", name, ", processed-", d)
		counter++
		if counter == 2 {
			builder.WriteString("Gorouting-")
			builder.WriteString(name)
			builder.WriteString(" Paniced")
			panic(builder.String())
		}

	}
}

func main() {
	var chanData chan int
	chanData = make(chan int, 5)

	var wg sync.WaitGroup

	wg.Add(1)
	go processData("First", &wg, chanData)
	wg.Add(1)
	go processDataPanic("Second (Panic)", &wg, chanData)

	for {
		sec := time.Now().Second()

		chanData <- sec

		if sec == 59 {
			fmt.Print("I am done writing to channel")
			close(chanData)
			break
		}

		time.Sleep(1 * time.Second)

	}

	wg.Wait()

}
