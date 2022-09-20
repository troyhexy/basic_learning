package control

//条件语句不需要使用圆括号将条件包含起来 ()；
//无论语句体内有几条语句，花括号 {} 都是必须存在的；
//左花括号 { 必须与 if 或者 else 处于同一行；
//在 if 之后，条件语句之前，可以添加变量初始化语句
//if 关键字后面的条件判断表达式的求值结果必须是布尔类型，即要么是 true，要么是 false：
func GetIf() {
	var a bool = false
	var b int = 1

	if b == 1 {
		println("测试if加括号")
	}

	if a {
		println("测试if不叫括号")
	}
	//条件语句不需要使用圆括号将条件包含起来 ();

	//if b  println("测试if")  //无论语句体内有几条语句，花括号 {} 都是必须存在的；

	//if a
	//{
	//	println("测试if")
	//} //左花括号 { 必须与 if 或者 else 处于同一行；

	if c := 1; c == 1 {
		println("测试ifj加表达式")
	}
	//在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 间隔，

	//if b {
	//
	//}
	//if 关键字后面的条件判断表达式的求值结果必须是布尔类型，即要么是 true，要么是 false：

}

//操作符是有优先级
func GetIfPriority() {

	//&&优先级大于！=
	a, b := false, true
	if a && b != true {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false")

	//个人建议，尽量使用括号，避免阅读代码对炒作操作优先级的判断负担
	if a && (b != true) {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false")

}

//if 快乐路径
//仅使用单分支控制结构；
//当布尔表达式求值为 false 时，也就是出现错误时，在单分支中快速返回；
//正常逻辑在代码布局上始终“靠左”，这样读者可以从上到下一眼看到该函数正常逻辑的全貌；
//函数执行到最后一行代表一种成功状态。
//尽量将命中率高的判断放在执行前面
func GetIfHappyPath() int {

	errorCondition1 := false
	errorCondition2 := false

	if errorCondition1 {
		// some error logic ... ...
		return 1
	}

	if errorCondition2 {
		// some error logic ... ...
		return 1
	}
	// some success logic
	return 2

}


