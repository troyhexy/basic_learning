package demo

import "fmt"

//包级变量只能使用带有 var 关键字的变量声明形式，不能使用短变量声明形式，但在形式细节上可以有一定灵活度
//m:=10 短变量声明形式 不通过

//var m int
//m=1   这样也无法编译通过

var m int =10  // var m =10 也可以

//声明但延迟初始化
var m1 int


func GetVar()  {



	//变量声明块（block）的语法形式
	var (
		a int = 128
		b int8 = 6
		s string = "hello"
		c rune = 'A'
		t bool = true
	)
	fmt.Println(a,b,s,c,t)

    //一行变量声明中同时声明多个变量
	var a1, b1, c1 int = 5, 6, 7
	fmt.Println(a1,b1,c1)

	// 结合上面二种方式
	var (
		a2, b2, c2 int = 5, 6, 7
		f2, d2, e3 rune = 'C', 'D', 'E'
	)

	fmt.Println(a2,b2,c2)
	fmt.Println(f2,d2,e3)


	//通用的变量声明的基础上，Go 编译器允许我们省略变量声明中的类型信息
	var a4  = 13
	fmt.Println(a4)
	 //可以使用这种变量声明“语法糖”声明多个不同类型的变量
	var a3, b3, c3 = 12, 'A', "hello"
	fmt.Println(a3,b3,c3)

	//当我们不接受默认类型
	var a5 int8 = 1
	var a6 =int8(1)
	fmt.Println(a5,a6)



}

//短变量声明
func GetVarShort()  {

	a := 12
	b := 'A'
	c := "hello"
	fmt.Println(a,b,c)

	a1, b1, c1 := 12, 'A', "hello"
	fmt.Println(a1,b1,c1)

//	a:=1 无法编译通过
	a,d:=test()
	fmt.Println(a,d) //短变量声明赋值，多个返回值接收变量只要有一个未声明，就可以覆盖
//	e,c:=test() //覆盖时必须是同类型


}

func test() (int ,int) {
	return 1, 1
}







func GetVarRule()  {

	//这两种方式都是可以使用的，但从声明一致性的角度出发
	//Go 更推荐我们使用第一个，这样能统一接受默认类型和显式指定类型这两种声明形式，
	//尤其是在将这些变量放在一个 var 块中声明时，你会更明显地看到这一点
	var (
		a = 13
		b = int32(17)
		f = float32(3.14)
	)

	fmt.Println(a,b,f)
	var (
		a1 = 13
		b1 int32 = 17
		f1 float32 = 3.14
	)
	fmt.Println(a1,b1,f1)

	//声明聚类
	var (
		a2 bool
		b2 bool
	)
	var (
		a3 int
		b3 int
	)
	fmt.Println(a2,b2,a3,b3)
	//就近原则 即也就是说我们尽可能在靠近第一次使用变量的位置声明这个变量
	var a5 =11
	fmt.Println(a5)

	//对于声明且显式初始化的局部变量，建议使用短变量声明形式

	a6 := int32(17)
	f6 := float32(3.14)
	s6 := []byte("hello, gopher!")

	fmt.Println(a6,f6,s6)

	//尽量在分支控制时使用短变量声明形式

	for i := len(s6); i > 0; i++ {

	}


}