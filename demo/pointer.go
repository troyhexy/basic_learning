package demo

import (
	"fmt"
	"unsafe"
)

func GetPointer() {

	var p *int
	fmt.Println(p) //输出 nil
	var a = 100
	p = &a
	fmt.Println(p) //输出 内存地址的值
	*p = 20        //修改指针值
	fmt.Println(a)

	//
	p1 := new(int)
	var b int
	p1 = &b
	*p1 = 100
	println(b)

}

func GetUnsafePointer() {

	//任何类型的指针都可以被转化为 unsafe.Pointer；
	//unsafe.Pointer 可以被转化为任何类型的指针；
	//uintptr 可以被转化为 unsafe.Pointer；
	//unsafe.Pointer 可以被转化为 uintptr。
	i := 10
	var p *int = &i
	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(i) // 100

	a := 10.00
	var fp1 *float64 = &a
	var fp2 *int = (*int)(unsafe.Pointer(fp1))
	*fp2 = *fp2 * 10
	fmt.Println(a) // -6.999477138774142e-302

	arr := [3]int{1, 2, 3}
	ap := &arr
	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	//unsafe.Pointer 这个桥梁转化为 uintptr 类型，再加上数组元素偏移量（通过 unsafe.Sizeof 函数获取），就可以得到该数组第二个元素的内存地址
	*sp += 3
	fmt.Println(arr) // [1 5 3]

}

//二级指针的使用
func GetSecondPointer() {

	//二级指针就可以用来改变指针变量的值，也就是指针变量的指向
	var a int = 5
	var p1 *int = &a
	println(*p1) // 5
	var b int = 55
	var p2 *int = &b
	println(*p2) // 55

	var pp **int = &p1
	println(**pp) // 5

	println((*pp) == p1) // true
	println((**pp) == (*p1)) // true
	println((**pp) == a)     // true

	pp = &p2
	println(**pp) // 55


}
