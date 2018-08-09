package main

import (
	"fmt"
	"sync"
)

//这个是完善原先并发的不足，并作出一些更改

//自己实现
/*func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
		go func() {
			done <- true
		}()  //这个是为了防止一组小写的打印完，大写的就卡住了，这个方法未必好，但是可以解决这个问题
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

//自己实现
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, w := range workers {
		w.in <- 'a' + i
		//<-workers[i].done  //这样就不是并行的
	}

	for i, w := range workers {
		w.in <- 'A' + i
		//<-workers[i].done
	}
	//都发完之后再打印，就会是并行的
	for _, w := range workers {
		<-w.done
		<-w.done
	}
}*/

//用go带的包实现
func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func chanDemo() {

	var wg *sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, wg)
	}

	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()

}
