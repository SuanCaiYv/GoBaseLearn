package src

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Chap2 /*
/*
这里指出了如何创建一个变量，以及之后怎么初始化它，还有"象鼻命名法"的使用
Go里有四种生命一个变量的方式：var, const, type, func
func声明一个函数变量，type声明一个新的类型，const声明一个常量
*/
func Chap2() {
	var val1 = 1
	var val2 int
	val2 = 2
	const val3 = 3
	type T1 int
	var val4 T1
	val4 = 4
	val5 := func() {
		fmt.Println("aaa")
	}
	val5()
	fmt.Println(val1, val2, val3, val4)
	x := 1
	// Go语言称为互联网时代的C语言不无道理，以为它也提供了指针操作，基本概念和使用和C很像
	xP := &x
	fmt.Println(*xP)
	// new()是一个内置函数，用来创建一个指定类型的空间，并返回内存地址
	// new()创建的空间不一定是在堆上，也可能在栈上，通过指针可以访问这个空间
	p1 := new(int)
	*p1 = 2
	// 虽然T2和T3都是int32类型，但是在编译时会视为不同类型，这样做的好处是通过编译器保证类型安全
	// 此外，可以使用T(x)来进行强制类型转换
	type T2 int32
	type T3 int32
}

func Chap3() {
	str1 := "aaa和我"
	// 返回字节数组
	byte1 := []byte(str1)
	// 返回码点数组
	rune1 := []rune(str1)
	fmt.Println(len(byte1), len(rune1))
	fmt.Println(strconv.ParseInt("123", 10, 32))
}

func Chap4() {
	// Go中的数组长度必须是常量，所以不是很方便，很多时候我们会用切片(可变数组)
	// 此外还有一个原因，就是当数组作为参数进行传递时，会触发整个数组的复制操作
	var arr1 [10]int
	arr1[0] = 1
	// 使用内置函数make()创建一个切片，并指出len和cap
	// 切片底层是数组，切片本身是一个结构体，这个结构体保存了底层数组的指针和长度，容量等信息
	slice1 := make([]int, 10, 20)
	slice1[0] = 1
	// 切片有很多类型，Go内置函数append()提供了向切片追加数据的功能，因为可能会触发底层数组的扩容
	// 所以此时append返回值有可能是新的slice
	slice1 = append(slice1, 1)
	map1 := make(map[string]int)
	map1["aaa"] = 1
	// 如果某个key不存在，则会得到value的零值，而不是nil
	if map1["bbb"] == 0 {
		fmt.Println("not exist")
	}
	// 可以通过多返回值判断是否存在key
	if _, ok := map1["ccc"]; ok {
		fmt.Println("existed")
	}
	// Go内置函数实现对map的key的删除操作
	delete(map1, "aaa")
}

// 稍微解释一下，第一个out()是匿名函数的名字，因为涉及到了闭包，所以外层函数最好有一个名字
// 此外，func()指出闭包的定义，既参数列表和返回值
func out() func(string, string) string {
	// 匿名函数的变量，所有闭包共享
	var init string
	fmt.Println("run0")
	// 所以这里的闭包的参数列表和返回值必须和上面的定义一一对应
	ans := func(s1, s2 string) string {
		fmt.Println("run")
		init += s1 + s2
		return init
	}
	fmt.Println("run1")
	return ans
}

type CanDo struct {
}

func Chap5() {
	// 匿名函数
	f1 := func(str string) {
		fmt.Println(str)
	}
	// 函数可以作为值，作为参数进行传递
	f1("aaa")
	f2 := out()
	f3 := out()
	str1 := f2("aaa", "bbb")
	fmt.Println(str1)
	str1 = f2("ccc", "ddd")
	// 因为是对同一个匿名函数的调用，所以共享函数内的局部变量
	// 最终的多次执行或转换成对于闭包的多次执行，既，同一个匿名函数只会在赋值时被执行一次，接下来都会变成闭包的执行
	// 所以会触发匿名函数的变量共享
	fmt.Println(str1)
	str1 = f3("eee", "fff")
	fmt.Println(str1)
	str1 = f3("ggg", "hhh")
	fmt.Println(str1)
	flag := rand.Int() % 2
	// defer通过压入堆栈的方式实现调用顺序反转，所以多个defer执行顺序和编写顺序相反
	defer func() {
		// recover()函数获取上下文的panic，然后我们进行自定义处理
		switch cause := recover(); cause {
		case nil:
		case CanDo{}:
			// 是我们想要处理的异常类型
			fmt.Println("just an exception")
		default:
			// 把panic传递下去
			panic(cause)
		}
	}()
	if flag == 0 {
		panic("an error")
	} else {
		panic(CanDo{})
	}
}
