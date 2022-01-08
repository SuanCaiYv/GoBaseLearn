package base

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var (
	// 定义一个互斥锁
	lock sync.Mutex
	// 临界区资源
	count int
)

var (
	// 定义一个读写锁
	rwLock sync.RWMutex
	// 临界区资源
	rwCount int
)

var (
	// 定义一个单例模式
	single sync.Once
	// 单例
	slice []int
)

func adder1() {
	lock.Lock()
	count += 1
	lock.Unlock()
}

func adder2() {
	lock.Lock()
	count += 1
	lock.Unlock()
}

func adder3() {
	lock.Lock()
	count += 1
	lock.Unlock()
}

func adder4() {
	rwLock.Lock()
	rwCount += 1
	rwLock.Unlock()
}

func adder5() {
	rwLock.Lock()
	rwCount += 1
	rwLock.Unlock()
}

func reader1() int {
	defer rwLock.RUnlock()
	rwLock.RLock()
	return rwCount
}

func reader2() int {
	defer rwLock.RUnlock()
	rwLock.RLock()
	return rwCount
}

func reader3() int {
	defer rwLock.RUnlock()
	rwLock.RLock()
	return rwCount
}

func reader4() []int {
	// Once#Do()只会被执行一次，底层通过CAS保证
	single.Do(func() {
		slice = make([]int, 10, 20)
		for i := 0; i < len(slice); i += 1 {
			slice[i] = i
		}
	})
	return slice
}

func fn91() {
	lock = sync.Mutex{}
	rwLock = sync.RWMutex{}
	single = sync.Once{}
	go func() {
		for i := 0; i < 1000; i += 1 {
			adder1()
		}
	}()
	go func() {
		for i := 0; i < 1000; i += 1 {
			adder2()
		}
	}()
	go func() {
		for i := 0; i < 1000; i += 1 {
			adder3()
		}
	}()
	go func() {
		for i := 0; i < 1000; i += 1 {
			adder4()
		}
	}()
	go func() {
		for i := 0; i < 1000; i += 1 {
			adder5()
		}
	}()
	go func() {
		for i := 0; i < 50; i += 1 {
			fmt.Printf("r1: %d\n", reader1())
		}
	}()
	go func() {
		for i := 0; i < 50; i += 1 {
			fmt.Printf("r2: %d\n", reader2())
		}
	}()
	go func() {
		for i := 0; i < 50; i += 1 {
			fmt.Printf("r3: %d\n", reader3())
		}
	}()
	go func() {
		for val := range reader4() {
			fmt.Printf("rr1: %d\n", val)
		}
	}()
	go func() {
		for val := range reader4() {
			fmt.Printf("rr2: %d\n", val)
		}
	}()
	go func() {
		for val := range reader4() {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Duration(math.MaxInt))
}

func ShowConcurrent() {
	fn91()
}
