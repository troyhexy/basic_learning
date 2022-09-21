package function

import "fmt"


var addv7 = func(a, b int) int {
	return a + b
}

func Test() {
	// 1、将匿名函数赋值给变量
	 addv8 := func(a, b int) int {
		return a + b
	}
	// 调用匿名函数 add
	fmt.Println(addv7(1, 2))
	fmt.Println(addv8(1, 2))
	// 2、定义时直接调用匿名函数
	func(a, b int) {
		fmt.Println(a + b)
	} (1, 2)

}

//匿名函数使用场景



//1保证局部变量的安全性
//匿名函数内部声明的局部变量无法从外部修改，从而确保了安全性
func UseAnonymousFunctionV1()  {
	var j int = 1
	f := func() {
		var i int = 1
		fmt.Printf("i, j: %d, %d\n", i, j)
	}
	f()
	j += 2
	f()
}


//将匿名函数作为函数参数
func UseAnonymousFunctionV2()  {
	add := func(a, b int) int {
		return a - b
	}
	// 将函数类型作为参数
	func(call func(int, int) int) {
		fmt.Println(call(1, 2))
	}(add)


}

//将匿名函数作为函数返回值

func UseAnonymousFunctionV3()  {
	// 此时返回的是匿名函数
	addFunc := deferAdd(1, 2)
	addFunc1 := deferAdd(1, 3)
	// 这里才会真正执行加法操作
	fmt.Println(addFunc())
	fmt.Println(addFunc1())


}

func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

