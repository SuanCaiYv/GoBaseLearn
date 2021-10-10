package src

import (
	"fmt"
	"time"
)

func compute1() {
	sum := 0
	for i := 0; i < 50; i += 1 {
		sum += i
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	fmt.Println("now is: " + time.Now().String())
}

func fn81() {
	fmt.Println("now: " + time.Now().String())
	// 使用Go的协程，去执行一个耗时应用。
	go compute1()
	fmt.Println("now: " + time.Now().String())
	time.Sleep(time.Duration(10) * time.Second)
}

// 通过在chan前后指明<-来指出通道的单向性，即只读或只写，对于只读的通道调用close()会出发panic
func consumer1(input <-chan string) {
	for {
		val, ok := <-input
		if !ok {
			break
		}
		fmt.Println("consumer1 get: " + val)
	}
}

func consumer2(input chan string) {
	for val := range input {
		fmt.Println("consumer2 get: " + val)
	}
}

func pipeline1() {
	// 指定通道缓冲区大小
	output := make(chan string, 5)
	another := make(chan string)
	go pipeline2(output, another)
	for i := 0; i < 10; i += 1 {
		output <- fmt.Sprintf("%08d", i)
	}
	close(output)
}

func pipeline2(input, output chan string) {
	go pipeline3(output)
	for val := range input {
		val = "pipeline2: " + val
		output <- val
	}
	close(output)
}

func pipeline3(input chan string) {
	for val := range input {
		fmt.Println(val)
	}
}

func fn82() {
	out1 := make(chan string)
	out2 := make(chan string)
	go consumer1(out1)
	go consumer2(out2)
	for i := 0; i < 10; i += 1 {
		out1 <- fmt.Sprintf("%04d", i)
	}
	for i := 0; i < 10; i += 1 {
		out2 <- fmt.Sprintf("%04d", 10-i)
	}
	close(out1)
	close(out2)
	pipeline1()
	time.Sleep(time.Duration(10) * time.Second)
}

func ShowGoroutine() {
	fn82()
}
