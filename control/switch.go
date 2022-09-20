package control

//不需要在每个分支结构中显式通过 break

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2() int {
	println("eval case2 expr")
	return 2
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}

func GetSwitch() {
	switch switchexpr() {
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}
	//查看case 执行顺序

	//输出
	//eval switch expr
	//eval case1 expr
	//eval case2_1 expr
	//eval case2_2 expr
	//exec case2

	//结论
	//Go 先对 switch expr 表达式进行求值，然后再按 case 语句的出现顺序，从上到下进行逐一求值
	//在带有表达式列表的 case 语句中，Go 会从左到右，对列表中的表达式进行求值，比如示例中的 case2_1 函数就执行于 case2_2 函数之前
	//如果 switch 表达式匹配到了某个 case 表达式，那么程序就会执行这个 case 对应的代码分支，比如示例中的“exec case2”。这个分支后面的 case 表达式将不会再得到求值机会
	//即便后面的 case 表达式求值后也能与 switch 表达式匹配上，Go 也不会继续去对这些表达式进行求值了。
	//无论 default 分支出现在什么位置，它都只会在所有 case 都没有匹配上的情况下才会被执行的。
	//考虑到 switch 语句是按照 case 出现的先后顺序对 case 表达式进行求值的，那么如果我们将匹配成功概率高的 case 表达式排在前面，就会有助于提升 switch 语句执行效率

	//switch 语句在 case 中支持表达式列表
	var a = 1
	switch a {
	case 1, 2, 3, 4, 5:
		println("it is a work day")
	case 6, 7:
		println("it is a weekend day")
	default:
		println("are you live on earth")
	}

	//fallthrough 用法
	switch switchexpr() {
	case case1():
		println("exec case1")
		fallthrough
	case case2():
		println("exec case2")
		fallthrough
	default:
		println("exec default")
	}
	//输出
	//eval switch expr
	//eval case1 expr
	//eval case2 expr
	//exec case2
	//exec default

}

type person struct {
	name string
	age  int
}

func GetSwitchType() {
	//只要类型支持比较操作，都可以作为 switch 语句中的表达式类型。
	//比如整型、布尔类型、字符串类型、复数类型、元素类型都是可比较类型的数组类型，
	//甚至字段类型都是可比较类型的结构体类型
	p := person{"tom", 13}
	switch p {
	case person{"tony", 33}:
		println("match tony")
	case person{"tom", 13}:
		println("match tom")
	case person{"lucy", 23}:
		println("match lucy")
	default:
		println("no match")
	}

	//switch 表达式的类型为布尔类型时，如果求值结果始终为 true，那么我们甚至可以省略 switch 后面的表达式
	switch {
	case false:
		println("match false")
	case true:
		println("match true")
	}

	//“type switch”这是一种特殊的 switch 语句用法，
	//switch 关键字后面跟着的表达式为x.(type)，这种表达式形式是 switch 语句专有的，而且也只能在 switch 语句中使用。
	//这个表达式中的 x 必须是一个接口类型变量，表达式的求值结果是这个接口类型变量对应的动态类型

	var x interface{} = 13
	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int")
	case string:
		println("the type of x is string")
	case bool:
		println("the type of x is string")
	default:
		println("don't support the type")
	}

	switch v := x.(type) {
	case nil:
		println("v is nil")
	case int:
		println("the type of v is int, v =", v)
	case string:
		println("the type of v is string, v =", v)
	case bool:
		println("the type of v is bool, v =", v)
	default:
		println("don't support the type")
	}


	//我们一直使用 interface{}这种接口类型的变量，
	//Go 中所有类型都实现了 interface{}类型，所以 case 后面可以是任意类型信息。
	//但如果在 switch 后面使用了某个特定的接口类型 I，那么 case 后面就只能使用实现了接口类型 I 的类型了

	//var t T
	//var i I = t
	//switch i.(type) {
	//case T:
	//	println("it is type T")
	//case int:
	//	println("it is type int")
	//case string:
	//	println("it is type string")
	//}

}

type I interface {
	M()
}
type T struct{}

func (T) M() {}
