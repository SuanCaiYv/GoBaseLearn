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
	go compute1()
	fmt.Println("now: " + time.Now().String())
	time.Sleep(time.Duration(10) * time.Second)
}

func ShowGoroutine() {
	fn81()
}