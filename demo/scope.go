package demo

import "fmt"

//作用域


var a1  = 11
func foo(n int) {
	a1 := 1
	a1 += n
}

func foo1(n int)  {
	var p *int =&a1
	*p+=n
}

func GetScope() {
	fmt.Println("a1 =", a1)
	foo(5)
	fmt.Println("after calling foo, a =", a1) // 11

	foo1(5)
	fmt.Println("after calling foo1, a =", a1) // 16

	bar()
}


//对于控制语句隐式代码块中的标识符的作用域划分
func bar() {
	if aa := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(aa, b, c)
	}
	//println(aa) 当前作用域未发现变量

	switch aa:=1;aa{
	case 2020:
		fmt.Println("it is", aa)
		bb:=21;
		fmt.Println("it is", bb)
	case 2021:
		fmt.Println("it is", aa)
		//fmt.Println("it is", bb)当前作用域未发现变量
	}
	//fmt.Println("it is", aa) 当前作用域未发现变量


	for i := 1; i <= 10; i++ {
		aa:=12
		fmt.Println("it is", aa)
	}

	//fmt.Println("it is", i) 当前作用域未发现变量
	//fmt.Println("it is", aa) //当前作用域未发现变量

	//bb:=11;
	//bb:=12;无法重复定义
	for i := 1; i <= 10; i++ {
		bb:=12
		fmt.Println("it is", bb)//for 相当与重复执行10个代码块，从而进行的变量覆盖即重复定义
	}


}

