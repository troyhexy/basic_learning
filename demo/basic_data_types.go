package demo

import (
	"fmt"
)

/**
*查看golang中涉及的基本类型
 **/

//数值整形
func GetInt() {

	//整型
	var a int8  //	8-bit integers (-128 to 127)
	var b int16 //	16-bit integers (-32768 to 32767)
	var c int32 //	32-bit integers (-2147483648 to 2147483647)
	var d int64 //	64-bit integers (-9223372036854775808 to 9223372036854775807)

	fmt.Println(a) //初始化后的默认值 0
	fmt.Println(b) //初始化后的默认值 0
	fmt.Println(c) //初始化后的默认值 0
	fmt.Println(d) //初始化后的默认值 0

	//无符号整形
	var ua uint8  //	8-bit integers (0 to 255)
	var ub uint16 //	16-bit integers (0 to 65535)
	var uc uint32 //	32-bit integers (0 to 4294967295)
	var ud uint64 //	64-bit integers (0 to 18446744073709551615)

	fmt.Println(ua) //初始化后的默认值 0
	fmt.Println(ub) //初始化后的默认值 0
	fmt.Println(uc) //初始化后的默认值 0
	fmt.Println(ud) //初始化后的默认值 0

	var i int       //32位的平台，为4字节，64位的平台为8字节
	var iu uint     //32位的平台，为4字节，64位的平台为8字节
	fmt.Println(i)  //初始化后的默认值 0
	fmt.Println(iu) //初始化后的默认值 0
	var p uintptr  //golang 对这个类型的描述，一个可以大到存储指针的值
	fmt.Println(p)
	//结论 int 整形初始化后默认赋值为0
}

//浮点类型
func GetFloat() {

	var a float32 //	32-bit integers (-3.403E38 to 3.403E38)
	var b float64 //	64-bit integers (-1.798E308 to 1.798E308)

	fmt.Println(a) //初始化后的默认值 0
	fmt.Println(b) //初始化后的默认值 0

}

//布尔型
func GetBool() {
	var a bool
	fmt.Println(a)
	//结论 布尔类型初始化后默认赋值为false ,取值范围 false,true
}

//字符类型
func GetByte() {
	var a byte                  //type byte = uint8
	fmt.Println(a)              //默认值为0
	fmt.Printf("a的值是：%c \n", a) //初始化 UTF-8值 0，对应指的值为 nut ,空格

	a = 'a'
	fmt.Println(a)
	fmt.Printf("a的值是：%c \n", a)

	//go中，字符的本质是一个整数，直接输出时，会输出它对应的UTF-8编码的值
	var b byte = 100

	fmt.Println(b)
	fmt.Println(a + b)
	fmt.Printf("a+b的值是：%c \n", a+b)

	//判断当前相加值，是否溢出uint8
	var c byte = 200
	fmt.Println(a + c)
	fmt.Printf("a+c的值是：%c \n", a+c)



}

//字符串类型
func GetString() {

	var a string
	fmt.Println(a) //默认值为空字符串
	fmt.Printf("a的值是：%s方便显示 \n", a)
	//a='abc' 这样编译不会通过，需要双引号
	a = "abd"
	a = "ahdwivh\n" //转义字符
	fmt.Println(a)
	//原样书输出
	a = `kslvf\n
           sdvsd`
	fmt.Println(a)

}

//复数类型
func GetComplex()  {
	var a complex64 //32 位实部和虚部）
	var b complex128 //64 位实部和虚部）
	fmt.Println(a)//默认(0+0i)
	fmt.Println(b)//默认(0+0i)
	a=1.1
	b=13i
	fmt.Println(a)
	fmt.Println(b)

}

//Unicode 字符
func GetRune() {
	var a rune                  //type rune = int32
	fmt.Println(a)              //默认值为0
	fmt.Printf("a的值是：%c \n", a) //初始化 UTF-8值 0，对应指的值为 nut ,空格
	a = '中'
	fmt.Println(a)
	fmt.Printf("a的值是：%c \n", a)
}


