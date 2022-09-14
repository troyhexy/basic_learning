package demo

import (
	"fmt"
)

//常量是指编译期间就已知且不可改变的值，常量只可以是数值类型（包括整型、 浮点型和复数类型）、布尔类型、字符串类型等标量类型

const Pi float64 = 3.14159265358979323846
const zero = 0.0 // 无类型浮点常量
const (          // 通过一个 const 关键字定义多个常量，和 var 类似
	size int64 = 1024
	eof        = -1 // 无类型整型常量
)
const u, v float32 = 0, 3   // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量

const b1 = b

//const w int
//w=1 编译无法通过

//const p  = &b 指针类型无法作为常量
//const ErrShortWrite = errors.New("short write")
//const s=[]int{1,2}
//Pi=5 常量编译期间无法修改
//const b=5 常量编译期间无法修改

//预定义常量
//Go 语言预定义了这些常量：true、false 和 iota。

//iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为 0，然后在下一个 const 出现之前，每出现一次 iota，其所代表的数字会自动增 1。

const holiday = iota
const holiday1 = iota

//const (
//	Sunday = iota
//	Monday = iota
//	Tuesday= iota
//	Wednesday= iota
//	Thursday= iota
//	Friday= iota
//	Saturday= iota
//	numberOfDays= iota
//)   可以简写为下面的形式

const (
	Sunday       = iota //0
	Monday              //1
	Tuesday             //2
	Wednesday           //3
	Thursday            //4
	Friday              //5
	Saturday            //6
	numberOfDays        //7
)

const (
	IgEggs         = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                // 1 << 1 which is 00000010
	IgNuts                     // 1 << 2 which is 00000100
	IgStrawberries             // 1 << 3 which is 00001000
	IgShellfish                // 1 << 4 which is 00010000
)

func GetConst() {
	fmt.Println(Pi)
	fmt.Println(b1)

	//Pi=4 编译无法通过



	//常量覆盖
	const b = 5
	const Pi = 5
	fmt.Println(b)
	fmt.Println(Pi)

	println(holiday)
	println(holiday1)

	println(Sunday)
	println(Monday)

	println(IgEggs)
	println(IgChocolate)
	println(IgShellfish)

}
