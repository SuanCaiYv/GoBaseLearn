package base

import (
	"fmt"
	"math/rand"
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
		// 当chan被关闭时，这里会得到一个通知，即ok为false，当然还有一种和for循环搭配的方式
		val, ok := <-input
		if !ok {
			break
		}
		fmt.Println("consumer1 get: " + val)
	}
}

func consumer2(input chan string) {
	// 这里的for循环是一种简化写法，编译之后后变成上面那种带ok检测的
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

func producer1(out chan int) {
	for i := 0; i < 10; i += 1 {
		t := rand.Int31() % 1000
		time.Sleep(time.Duration(t) * time.Millisecond)
		out <- i
	}
	close(out)
}

func producer2(out chan int) {
	for i := 10; i < 20; i += 1 {
		t := rand.Int31() % 1000
		time.Sleep(time.Duration(t) * time.Millisecond)
		out <- i
	}
	close(out)
}

func fn83() {
	input1 := make(chan int)
	input2 := make(chan int)
	go producer1(input1)
	go producer2(input2)
	count := 0
	// select有一些坑，我说一下：
	// 首先，select在有多个可用的case时，会在当次循环随机选择一个，而不是全部执行，剩余'就绪的'会在下一次循环执行
	// 刚刚说到了循环，这里说一下，select在执行一次之后，不会再次执行，所以我们常写在循环中，做到多次调用select
	// 最后，select支持default语句，也就是在没有任何case就绪时执行的选择，所以一般来说for循环不会和default搭配使用，否则会一直执行default
	// 如果没有default，select会被阻塞住，即不运行，直到有新的case就绪为止
	// 现在我们来分类讨论一下：
	// 首先是有就绪case：1⃣️随机选择一个case执行
	// 没有就绪case：1⃣️如果没有default语句，则阻塞；2⃣️如果有default语句，则执行default语句
	for {
		// select多路复用技术，对多个chan进行监听，有可用的则执行
		select {
		// 这里也有坑，如果chan被关闭，case会读取chan指定数据类型的零值
		// 所以如果只是单纯地读取，然后进行零值判断可能会误判(人家就真的发了个0呢？)
		// 所以引入ok判断，然后置空，因为select会跳过所有chan为nil的case
		case val, ok := <-input1:
			if !ok {
				count += 1
				input1 = nil
			} else {
				fmt.Printf("get1: %d\n", val)
			}
			// 如果close某一个chan之后，再写入会报错；如果试图读取一个已经关闭的chan，会读取到零值
			// 因为读取close之后的chan会读取到零值，所以我们更可能用ok检测法来检测，并定义一个全新的，专门用于执行退出操作的chan
			// 或者干脆向done这个chan写数据表明希望退出
		case val, ok := <-input2:
			if !ok {
				count += 1
				input2 = nil
			} else {
				fmt.Printf("get2: %d\n", val)
			}
		}
		if count == 2 {
			break
		}
	}
}

func ShowGoroutine() {
	fn83()
}
