package main

import (
	"fmt"
	"time"
)

type MyInterface interface {
	Foo()
}

type MyType struct {
}

func (m *MyType) Foo() {
}

func main() {
	closure1()
	time.Sleep(1 * time.Second)
	closure2()
	time.Sleep(1 * time.Second)
	closure3()
	time.Sleep(1 * time.Second)
}

func closure1() {
	work := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			<-work
			fmt.Println("1: ", i)
		}(i)
	}
	close(work)
}

func closure2() {
	work := make(chan struct{})
	for i := 0; i < 10; i++ {
		j := i
		go func() {
			<-work
			fmt.Println("2: ", j)
		}()
	}
	close(work)
}

func closure3() {
	work := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			<-work
			fmt.Println("3: ", i)
		}()
	}
	close(work)
}
