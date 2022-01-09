package c_go

/*
#include <stdio.h>
// 导入自定义头文件
#include "c_go_t1.h"

// 直接在注释实现
void print_str_v0(const char *str) {
	printf("%s\n", str);
}

void print_str_v1(const char *str);
*/
import "C"

func CGo1() {
	C.print_str_v0(C.CString("直接使用.go文件定义的C函数"))
	C.print_str_v1(C.CString("直接使用.c文件定义的函数"))
	C.print_str_v2(C.CString("使用头文件定义的函数，并由C实现"))
}
