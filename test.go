package main

import (
	"time"
	"fmt"
	"sync"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10] chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond) //给定一个时间，让其打印完毕，这种做法是不好的
}

var wg sync.WaitGroup

func test(show int) {
	fmt.Println(show)
	wg.Done()

}

func main() {
	//chanDemo()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("done!")
}
