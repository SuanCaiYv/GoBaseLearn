package cgo

/*
#include <errno.h>

// 这是一个，通过C的errno实现返回error的函数
double div(int a, int b) {
    if (b == 0) {
        errno = EINVAL;
        return 0;
    }
    return 1.0 * a / 1.0 * b;
}
*/
import "C"
import "fmt"

func CGo3() {
	v, err := C.div(1, 0)
	fmt.Println(v, err)
	v, err = C.div(3, 2)
	fmt.Println(v, err)
	// 此外，返回值为void的C函数，对应Go的_Ctype_void类型的返回值，即，即使C函数为void，Go也是有返回值的。

	// 关于C调用Go实现的函数，前面已经演示过了，这里只说一点，就是当前文件如果想立即使用导出C函数的Go实现体
	// 那可能做不到，因为存在循环依赖问题，即此时_cgo_export.h还未创建，但是它的创建依赖当前
}
