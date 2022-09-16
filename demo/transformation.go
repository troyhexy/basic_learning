package demo

import "strconv"

//基本数据转化
//GO 不支持隐式转换
//原类型和别名类型 也不可以
type MyInt int32

func GetImplicit() {
	var a int32 = 1
	var c int = 1
	var b int64
	//	b=a  也是无法向上兼容
	//	b=c 即使当前环境下是64位 也是无法编译的
	//	b=a

	var d MyInt
	//	d=a //别名无法通过
	println(a,b,c,d)

}

//数值类型的转换
func GetIntChange()  {

	var(
		a int8 = 1
		b int32 =2
		c float32 =1.8
		g rune = 'A'
		d int =300
		e string ="30000000000000000000000000000000000000000000000000"
	)


	 b= int32(a) //显示转换
	 a=int8(d)   //注意转换溢出
	 a=int8(c)   //浮点转整数舍弃小数位
	 a=int8(g)


	// 字符串转int
	d, err := strconv.Atoi(e) //转换字符串 为int
	f, err1 := strconv.ParseInt(e, 10, 8)//转换int64


	println(f,d,b,err,err1)
}


//字符串类型转换
func GetStringChange()  {

}

