package cgo

/*
#include "c_go_t2.h"
*/
import "C"

import "fmt"

func CGo2() {
	C.print_str_v3(C.CString("除了使用C/C++实现头文件定义的函数，还可以使用Go来实现"))
}

//export print_str_v3
func print_str_v3(str *C.char) {
	fmt.Println(C.GoString(str))
}
