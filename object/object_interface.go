package object

import "fmt"

type Interface1 interface {
	M1()
	//1Go 语言要求接口类型声明中的方法必须是具名的，并且方法名字在这个接口类型的方法集合中是唯一的
	//	M1(int)

}
type Interface2 interface {
	M1(string)
	M2()
}

//Go 1.14 版本以后，Go 接口类型允许嵌入的不同接口类型的方法集合存在交集，
//但前提是交集中的方法不仅名字要一样，它的方法签名部分也要保持一致，
//也就是参数列表与返回值列表也要相同，否则 Go 编译器照样会报错。
type Interface3 interface {
	Interface1
	//   Interface2 // 编译器报错：duplicate method M1
	M3()
}

func GetEmptyInterface() {
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
	v3 := i.(int64)
	fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
	v4 := i.([]int)                                     // panic: interface conversion: interface {} is int64, not []int
	fmt.Printf("the type of v4 is %T\n", v4)
}
