package unsafe

import (
	"fmt"
	"unsafe"
)

type Temp struct {
	a float64
	b int64
	c bool
}

func Unsafe() {
	t := Temp{
		a: 24.12,
		b: 12,
		c: true,
	}
	// 强制类型转换不是赋值，是原子性，啊也不对，强转本质只是解释方式改变，数据不变，所以一步到位
	pointerOfB := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + 8))
	*pointerOfB += 6
	fmt.Println(t)
	// 这里因为存在中间变量赋值，假如，就在赋值刚刚完成之后发生了GC导致t地址变化，那么tempAddress是不会被更新的
	tempAddress := uintptr(unsafe.Pointer(&t)) + 8
	// 所以会导致这里出现非法地址访问
	// https://stackoverflow.com/questions/42067478/when-is-it-safe-in-go-to-reference-an-object-only-through-a-uintptr有解释
	badPointerOfB := (*int64)(unsafe.Pointer(tempAddress))
	*badPointerOfB += 6
	fmt.Println(t)
}
