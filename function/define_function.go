package function

import (
	"errors"
	"fmt"
)

//在同一个 Go 包中，函数名应该是唯一的，
//并且它也遵守 Go 标识符的导出规则，也就是我们之前说的，首字母大写的函数名指代的函数是可以在包外使用的，小写的就只在包内可见。
func add(a int, b int) int {
	return a + b
}

//函数名称区分大小写
func Add(a int, b int) int {
	return a - b
}

//返回类型 单反值 可以使用 括号包裹，也可以不用
func addV1(a int, b int) int {
	return a + b
}

//多返回值，则必须使用括号
func addV2(a int, b int) (int, int) {
	return a, b
}

//入参类型相同，可以合并只写一个
func addV3(a, b, c int, d int8) int {
	return a + b + c + int(d)
}

//变长参数
//接受任意数量的参数，这些参数的类型全部是 int
func addV4(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

//具名返回
//1.必须使用 括号包裹返回值
//2.即使具名，函数体必须由return
func addV5(a int, b int) (c int) {

	c = a + b
	return

}

func addV6(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
}


