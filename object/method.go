package object

import (
	"fmt"
)

//注意点1• 只能为当前包内命名类型定义方法。
   //第一个推论：我们不能为原生类型（诸如 int、float64、map 等）添加方法。
   //第二个推论：不能跨越 Go 包为其他包的类型声明新方法。
//注意点2• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
//注意点3• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
//注意点4• 不支持方法重载，receiver 只是参数签名的组成部分。
//注意点5• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

//func (recevier type) methodName(参数列表)(返回值列表){}
//参数和返回值可以省略


//注意点3
//type MyInt *int
//func (r MyInt) String() string { // r的基类型为MyInt，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
//	return fmt.Sprintf("%d", *(*int)(r))
//}
//
//type MyReader io.Reader
//func (r MyReader) Read(p []byte) (int, error) { // r的基类型为MyReader，编译器报错：invalid receiver type MyReader (MyReader is an interface type)
//	return r.Read(p)
//}




type Test struct{}

// 无参数、无返回值
func (t Test) method0() {
	fmt.Println("method0")
}

//我们没有用到 receiver 参数，我们也可以省略 receiver 的参数名 即为注意点2
func (Test) method0v1() {
	fmt.Println("method0v1")
}

// 单参数、无返回值
func (t Test) method1(i int) {

}

// 多参数、无返回值
func (t Test) method2(x, y int) {

}

// 无参数、单返回值
func (t Test) method3() (i int) {
	return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (t *Test) method5() {
	fmt.Println("method5")
}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
	return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
	return
}

func GetMethod()  {
	 test :=Test{}
	 test1 :=&Test{}

	 //注意点5
	 test.method5()
	 test1.method0()
}


//Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
//两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
func GetMethodExpression() {

	// method value
	//隐式传递 receiver
	var test Test
	m := test.method0v1
	m()

	//method expression
	mExpression := (*Test).method0v1
	mExpression(&test) // 显式传递 receiver

}
