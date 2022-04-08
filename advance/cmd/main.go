package main

type MyInterface interface {
	Foo()
}

type MyType struct {
}

func (m *MyType) Foo() {
}

func main() {
	var i MyInterface
	i = &MyType{}
	i.Foo()
}
