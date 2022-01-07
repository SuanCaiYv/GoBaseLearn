package main

import (
	"fmt"
	"time"
)

/*
之前写了一份base版放在另一台电脑里了，现在重新写一份简略版的基础语法
*/

func main() {
	timeout := make(chan struct{}, 1)
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- struct{}{}
	}()
	<-timeout
	fmt.Println("run")
}
