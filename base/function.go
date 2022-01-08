package base

import (
	"errors"
	"fmt"
)

// Go通过栈来保存返回值，以此来实现多返回值，以及返回局部变量的可行性
// 不过这样无法使用寄存器来做返回值优化，因为只有一个返回值寄存器
func fn51(flag bool) (int, error) {
	if flag {
		return 0, nil
	} else {
		return -1, errors.New("err")
	}
}

func tmp1(a int) int {
	return a
}

func tmp2(a, b int) int {
	return a * b
}

func tmp3(a, b int) int {
	return a + b
}

func tmp4() (int, int) {
	return 1, 0
}

func fn2() {
	var f1 = tmp1
	var f2 = tmp2
	var f3 = tmp3
	var f4 = tmp4
	f1(1)
	// f2和f3是同一类型的函数值，他们值(函数定义)不同，但是类型(参数列表+返回值)相同
	f2(1, 2)
	f3(3, 4)
	f4()
}

// 注意，add()返回一个函数，而不是int
func add() func(a, b int) int {
	var sum = 0
	// 定义一个匿名函数
	var fn = func(a, b int) int {
		sum += a + b
		return sum
	}
	// 返回这个匿名函数作为add()的返回值
	return fn
}

func fn3() {
	var f1 = func(a, b int) int {
		return a + b
	}
	// 一个简单的使用匿名函数的例子，匿名函数允许在块里声明函数，它和普通函数的区别在于，没有函数名
	f1(1, 2)
	var sum = 0
	var f2 = func(a int) int {
		sum += a
		return sum
	}
	f2(1)
	f2(2)
	f2(3)
	// 匿名函数会捕获上下文环境
	fmt.Println(sum)
	// 函数值不仅仅拥有计算过程，还拥有状态
	var f3 = add()
	ans := f3(1, 2)
	fmt.Println(ans)
	// 匿名函数f3()拥有外部函数的变量的引用，所以拥有在它作用域之外的上下文
	// 或者可以认为，add()返回一个匿名函数的引用，每次对f3()的调用都会转换到对add()内部的匿名函数的调用
	// 所以sum才会出现累加的结果
	ans = f3(3, 4)
	fmt.Println(ans)
}

type markerType struct {
}

func fn4(flag int) {
	// 使用defer做收尾工作，它的实现是通过压入栈来实现的，所以多的defer的执行顺序和声明顺序相反
	defer func() {
		// 这里使用recover()获取上下文的panic
		switch cause := recover(); cause {
		case nil:
			// 根据panic种类进行特例化处理
		case markerType{}:
			// 这个panic是我们自定义异常，所以在这里处理
			fmt.Println("get an exception")
		default:
			// 其他的panic直接抛出即可
			panic(cause)
		}
	}()
	if flag == 0 {
	} else if flag < 0 {
		panic("err")
	} else {
		panic(markerType{})
	}
}

func ShowFunc() {
	fn3()
	fn4(0)
	fn4(1)
	fn4(-1)
}
