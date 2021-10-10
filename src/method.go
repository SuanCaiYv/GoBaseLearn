package src

import "fmt"

type complex0 struct {
	a int
	i int
}

func (c complex0) add1(c2 *complex0) *complex0 {
	return &complex0{
		a: c.a + c2.a,
		i: c.i + c2.i,
	}
}

func (c *complex0) add2(c2 *complex0) *complex0 {
	return &complex0{
		a: c.a + c2.a,
		i: c.i + c2.i,
	}
}

func (c *complex0) display() {
	if c == nil {
		fmt.Println("empty")
	} else {
		fmt.Printf("{a: %d, i: %d}\n", c.a, c.i)
	}
}

type point struct {
	x int
	y int
}

func (p *point) display() {
	fmt.Printf("{x: %d, i: %d}\n", p.x, p.y)
}

type node struct {
	point
	string
}

func addComplex(body func(c *complex0) *complex0, c2 *complex0) {
	ans := body(c2)
	ans.display()
}

func fn61() {
	var c1 = complex0{
		a: 1,
		i: 1,
	}
	var c2 = complex0{
		a: 2,
		i: 3,
	}
	c3 := c1.add2(&c2)
	c3.display()
	node1 := &node{
		point{
			x: 1,
			y: 2,
		},
		"aaa",
	}
	node1.display()
	// 方法值的使用，方法本身就是一个包含了接收者状态的函数，所以也可以作为一个值传递
	// 当它作为值传递时，也一同传递了接收者的状态
	addComplex(c1.add2, &c2)
	// 通过T.method()或(*T).method()可以进行方法查找，然后此时第一个参数会作为方法接收者传递，后面的参数依次和方法定义的参数列表一致
	(*complex0).add2(&c1, &c2).display()
}

func ShowMethod() {
	fn61()
}
