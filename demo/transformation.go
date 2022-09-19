package demo

import (
	"fmt"
	"strconv"
)

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
	var(
		a int8 = 1
		b float32 =1.8
		c rune = 'A'
		d bool =false
		e int =5
		f int64 =1
	)

	//int 转 string
	str := strconv.Itoa(e)

	//int64 转string
	str1 := strconv.FormatInt(f, 10)

	println(str,str1)

	// 转字符串
	str2 := fmt.Sprintf("%d", a)
	str3 := fmt.Sprintf("%f", b)
	str4 := string(c)
	str5 := fmt.Sprintf("%t", d)
	str6 := fmt.Sprintf("%d", e)

	println(str2,str3,str4,str5,str6)


}

func GetFloatChange()  {
	var(
		a int8 = 1
		b float32 =1.8
		s string ="1.5"
	)

	 f :=float32(a)
	 f2 :=float64(b)

	 //string 转 float64
	 f3,err := strconv.ParseFloat(s,32)

	println(f,f2,f3,err)


}

