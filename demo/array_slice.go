package demo

import "fmt"

//指针（pointer）


//通道（chan）
//结构体（struct）
//接口（interface）

//数组（array）
func GetArray() {

	var a [3]int
	fmt.Println(a)                                   //输出 [0,0,0]
	fmt.Printf("a的长度为:%d，容量为:%d \n", len(a), cap(a)) //默认长度容量为3
	//定义数组的几种方式
	var b = [3]int{1, 2, 3}
	var c = [...]int{1, 2, 3}
	var d = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}

	var f = new([3]int) //new(T)返回了一个指针，指向新分配的类型 T 的零值

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

}

//切片（slice）

func GetSlice() {

	var a []int
	fmt.Println(a)                                   //输出 []
	fmt.Printf("a的长度为:%d，容量为:%d \n", len(a), cap(a)) //默认长度容量为0

	if a==nil { //仅仅申明其 与nil 相等
		fmt.Printf("a与nil 比较\n")
	}



	var b = make([]int, 5, 6)
	fmt.Println(b)                                   //输出 [0,0,0,0,0]
	fmt.Printf("b的长度为:%d，容量为:%d \n", len(b), cap(b)) //默认长度5,容量6

	c := []int{1, 2, 3}
	fmt.Println(c)                                   //输出 [0,0,0]
	fmt.Printf("c的长度为:%d，容量为:%d \n", len(b), cap(b)) //默认长度3,容量3

}

func GetSliceByArr() {

	//基于数组创造切片
	arr := [...]int{1, 2, 3, 4, 5}
	b := arr[2:4]
	fmt.Println(b)                                   //[3,4]
	fmt.Printf("b的长度为:%d，容量为:%d \n", len(b), cap(b)) //默认长度2,容量为3
	//这里的长度为 4-2 =2，容量为 原来数组容量 5-2 =3

	//设置容量
	c := arr[2:4:4]
	fmt.Println(c)                                   //[3,4]
	fmt.Printf("c的长度为:%d，容量为:%d \n", len(c), cap(c)) //默认长度2,容量为2
	//这里的长度为 4-2 =2，容量为 我们设置的 4-2 =2
	//  c := arr[2:4:3]//编译不会通过，设置容量，要大于等于长度
	// c := arr[2:4:6]//编译不会通过，设置容量，不能超出原来容量5

	d := arr[:]
	fmt.Println(d)                                   //[1,2,3,4,5]
	fmt.Printf("d的长度为:%d，容量为:%d \n", len(d), cap(d)) //默认长度5,容量为5

	f := arr[:3]
	fmt.Println(f)                                   //[1,2,3]
	fmt.Printf("f的长度为:%d，容量为:%d \n", len(f), cap(f)) //默认长度3,容量为5

	g := arr[3:]
	fmt.Println(g)                                   //[4,5]
	fmt.Printf("g的长度为:%d，容量为:%d \n", len(g), cap(g)) //默认长度2,容量为2

}

//切片底层，解析
func GetSliceChange() {
	arr := [...]int{1, 2, 3, 4, 5}
	b := arr[2:4]
	c := arr[2:4:4]

	//基于数组的的创建，底层数据都引用数组,当数组修改时，会导致切边的改动
	arr[3] = 10
	fmt.Println(b) //[3,10]
	fmt.Println(c) //[3,10]

	//当切片扩容，重新生成数组，则与数组不影响
	b = append(b, 11)
	c = append(c, 11)
	arr[3] = 9
	fmt.Println(b) //[3,9,11]
	fmt.Println(c) //[3,10,11] //由于c 的容量为2，发生扩容后，底层数组指向改变，所以不影响

	//同理当切片未扩容时，也会影响底层数组
	arr1 := [...]int{1, 2, 3, 4, 5}
	d := arr1[2:4:5]
	fmt.Println(d) //[3,4]
	d = append(d, 99)
	fmt.Println(d)    //[3,4,99]
	fmt.Println(arr1) //[1,2,3,4,99],未扩容，等于修改最后数组arr1[4]

	//切片影响数组
	arr2 := [...]int{1, 2, 3, 4, 5}
	e := arr2[2:4:4]
	fmt.Println(e) //[3,4]
	e[0] = 12
	fmt.Println(e)    //[12,4]
	fmt.Println(arr2) //[1,2,12,4,5]

}

func GetSliceBySlice() {
	//基于切片创造切片
	var slice = []int{1, 2, 3, 4, 5, 6}

	b := slice[2:4]
	fmt.Printf("b的长度为:%d，容量为:%d \n", len(b), cap(b)) //默认长度2,容量为4
	c := slice[2:4:4]
	fmt.Printf("c的长度为:%d，容量为:%d \n", len(c), cap(c)) //默认长度2,容量为2
	d := slice[:]
	fmt.Printf("d的长度为:%d，容量为:%d \n", len(d), cap(d)) //默认长度6,容量为6
	f := slice[:3]
	fmt.Printf("f的长度为:%d，容量为:%d \n", len(f), cap(f)) //默认长度3,容量为6
	g := slice[3:]
	fmt.Printf("g的长度为:%d，容量为:%d \n", len(g), cap(g)) //默认长度3,容量为3

}
