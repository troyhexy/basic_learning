package exceptions

import (
	"fmt"
	"os"
)

//defer 执行顺序
//defer会在函数返回之前执行
//defer可以放在函数中任意位置 return 之前
//在 Go 中，只有在函数（和方法）内部才能使用 defer；
func GetDefer() int {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(5)
	defer fmt.Println(4)

	return 1
}

func GetDeferV2() {
	fmt.Println("A")
	defer func() {
		fmt.Println("B")
		defer fmt.Println("C")
		fmt.Println("D")
	}()

	defer fmt.Println("E")
	fmt.Println("F")

	//上面程序将依次输出： A F E B D C
}

//在调用os.Exit时候，defer不会执行。下面程序只会输出B
func GetDeferV3() {
	defer fmt.Println("A")
	fmt.Println("B")
	os.Exit(0)
}

//defer类似闭包使用时候，访问的总是循环中最后一个值
func DeferTrap() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	//修复方法，传入最后一个值
	for j := 0; j < 5; j++ {
		defer func(i int) {
			fmt.Println(j)
		}(j)
	}
}

//可以修改函数中的命名返回值

func DeferTrapV1() {
	fmt.Println(test())   //101
	fmt.Println(test2())  //1
	fmt.Println(*test3()) //101

	//返回值 = xxx
	//调用defer函数
	//空的return
	fmt.Println(f())  //1
	fmt.Println(f1()) //5
	fmt.Println(f2()) //1
}

func test() (i int) {
	defer func() {
		i++
	}()
	return 100
}

func test2() int {
	ret := 1
	defer func() {
		ret += 100
	}()
	return ret
}

func test3() *int {
	ret := 1
	defer func() {
		ret += 100
	}()
	return &ret
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//defer语句传参中，实参是函数语句，此时，会先执行函数，并将函数的返回值压入栈中，如：
func DeferTrapV2() {
	f := func(num int) int {
		fmt.Printf("%d ", num)
		return 10 * num
	}

	for index := 0; index < 5; index++ {
		defer fmt.Printf("%d ", f(index)) // 第一循环时，先执行函数f(0)，然后，将函数返回值入栈
	}
	fmt.Println("1111")
}

// defer1.go

func bar() (int, int) {
	return 1, 2
}

//明确哪些函数可以作为 deferred 函数
//从这组错误提示中我们可以看到，append、cap、len、make、new、imag 等内置函数都是不能直接作为 deferred 函数的，
//而 close、copy、delete、print、recover 等内置函数则可以直接被 defer 设置为 deferred 函数。
func foo() {
	var c chan int
	var sl []int
	var m = make(map[string]int, 10)
	m["item1"] = 1
	m["item2"] = 2
	var a = complex(1.0, -1.4)

	var sl1 []int
	defer bar()
	//	defer append(sl, 11)
	//	defer cap(sl)
	defer close(c)
	//	defer complex(2, -2)
	defer copy(sl1, sl)
	defer delete(m, "item2")
	//defer imag(a)
	//defer len(sl)
	//defer make([]int, 10)
	//defer new(*int)
	defer panic(1)
	defer print("hello, defer\n")
	defer println("hello, defer")
	//	defer real(a)
	defer recover()

	//可以使用闭包函数来满足要求
	defer func() {
		_ = real(a)
	}()
}
