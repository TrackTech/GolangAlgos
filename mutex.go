package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func processLineInner(line string) {
	words := strings.Split(line, " ")
	mutext.Lock()
	for _, word := range words {
		wordCount[word] += 1
		if wordCount[word] > maxCount {
			maxCount = wordCount[word]
			maxWord = word
		}
	}
	mutext.Unlock()
}

func processLine(task *sync.WaitGroup, line string) {
	defer task.Done()
	processLineInner(line)
}

var wordCount map[string]int
var maxCount int
var maxWord string

var mutext sync.Mutex

func main() {
	wordCount = map[string]int{}
	maxCount = 0

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var waitGroup sync.WaitGroup

	threadCount := 3
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			waitGroup.Wait()
			break
		}
		if threadCount > 0 {
			waitGroup.Add(1)
			go processLine(&waitGroup, string(line))
		} else {
			waitGroup.Wait()
			threadCount = 0
		}
	}
	fmt.Println(wordCount)
	fmt.Println("Max count=", maxCount, ", max word=", maxWord)

}
