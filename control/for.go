package control

import (
	"fmt"
	"time"
)

//和条件语句、分支语句一样，左花括号 { 必须与 for 处于同一行；
//不支持 whie 和 do-while 结构的循环语句；
//可以通过 for-range 结构对可迭代集合进行遍历；
//支持基于条件判断进行循环迭代；
//允许在循环条件中定义和初始化变量，且支持多重赋值；
//Go 语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break
func GetFor() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//Go 语言的 for 循环支持声明多循环变量，并且可以应用在循环体以及判断条件中
	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += (i + j + k)
		println(sum)
	}

	//环变量的更新操作放在了循环体中
	for i := 0; i < 10; {
		i++
	}

	//可以省略循环前置语句
	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	//也可以都省略
	i2 := 0
	for i2 < 10 {
		println(i2)
		i2++
	}
	//虽然我们对前置语句或后置语句进行了省略，但经典 for 循环形式中的分号依然被保留着，你要注意这一点，这是 Go 语法的要求

	//例外，那就是当循环前置与后置语句都省略掉，仅保留循环判断条件表达式时，我们可以省略经典 for 循环形式中的分号
	i3 := 0
	for i3 < 10 {
		println(i3)
		i3++
	}

	//当 for 循环语句的循环判断条件表达式的求值结果始终为 true 时，我们就可以将它省略掉了：
	for {
		// 循环体代码
	}

	for true {
		// 循环体代码
	}

	for {
		// 循环体代码
	}
}

func GetForRange() {

	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d] = %d\n", i, sl[i])
	}

	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}

	//当我们不关心元素的值时，我们可以省略代表元素值的变量 v，只声明代表下标值的变量 i：
	for i := range sl {
		// ...
		println(i)
	}

	//如果我们不关心元素下标，只关心元素值，那么我们可以用空标识符替代代表下标值的变量 i。这里一定要注意，这个空标识符不能省略
	for _, v := range sl {
		// ...
		println(v)
	}

	//我们既不关心下标值，也不关心元素值
	for _, _ = range sl {
		// ...
	}

	for range sl {
		// ...
	}
}

func GetForRangeStringMap() {
	//每次循环得到的 v 值是一个 Unicode 字符码点，也就是 rune 类型值，而不是一个字节，返回的第一个值 i 为该 Unicode 字符码点的内存编码（UTF-8）的第一个字节在字符串内存序列中的位置。
	var s = "中国人"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
	//我们要对 map 进行循环操作，for range 是唯一的方法，

	var m = map[string]int{
		"Rob":  67,
		"Russ": 39,
		"John": 29,
	}

	for k, v := range m {
		println(k, v)
	}

	//for range 会尝试从 channel 中读取数据
	//for range 每次从 channel 中读取一个元素后，
	//会把它赋值给循环变量 v，并进入循环体。
	//当 channel 中没有数据可读的时候，
	//for range 循环会阻塞在对 channel 的读操作上。
	//直到 channel 关闭时，
	//for range 循环才会结束，这也是 for range 循环与 channel 配合时隐含的循环判断条件
	var c = make(chan int)
	for v := range c {
		// ...
		println(v)
	}

}

func GetForContinue() {

	var sum int
	var sl1 = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(sl1); i++ {
		if sl1[i]%2 == 0 {
			// 忽略切片中值为偶数的元素
			continue
		}
		sum += sl1[i]
	}
	println(sum) // 9

	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 13, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == 13 {
				//发现每行第一个值为13 的位置
				fmt.Printf("found 13 at [%d, %d]\n", i, j)
				continue outerloop
			}
		}
	}

}

func GetForBreak() {

	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	// 找出整型切片sl中的第一个偶数
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			firstEven = sl[i]
			break
		}
	}

	println(firstEven) // 6

	var gold = 38

	var s = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 38, 127},
		{54, 27, 40, 83, 81},
	}
outerloop:
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == gold {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break outerloop
			}
		}
	}

}

func GetForBugs() {
	//循环变量的重用
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)

	//参数与 i、v 的当时值进行绑定，看下面的修正代码
	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 10)

	//参与循环的是 range 表达式的副本
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)

	//使用切片 修正
	for i, v := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)


	//break 跳出问题
	//不带 label 的 break 语句中断执行并跳出的，是同一函数内 break 语句所在的最内层的 for、switch 或 select
	var sl3 = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1


	// find first even number of the interger slice
	for i := 0; i < len(sl3); i++ {
		switch sl3[i] % 2 {
		case 0:
			firstEven = sl3[i]
			break
		case 1:
			// do nothing
		}
	}
	println(firstEven)



}
