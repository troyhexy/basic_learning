package demo

import (
	"fmt"
)

/*
查看golang中涉及的数值数据类型
 */
func GetInt() {

	//整型
	var a int8 //	8-bit integers (-128 to 127)
	var b int16//	16-bit integers (-32768 to 32767)
	var c int32//	32-bit integers (-2147483648 to 2147483647)
	var d int64//	64-bit integers (-9223372036854775808 to 9223372036854775807)


	fmt.Println(a) //初始化后的默认值 0
	fmt.Println(b) //初始化后的默认值 0
	fmt.Println(c) //初始化后的默认值 0
	fmt.Println(d) //初始化后的默认值 0


	//无符号整形
	var ua uint8    //	8-bit integers (0 to 255)
	var ub uint16  //	16-bit integers (0 to 65535)
	var uc uint32  //	32-bit integers (0 to 4294967295)
	var ud uint64  //	64-bit integers (0 to 18446744073709551615)

	fmt.Println(ua) //初始化后的默认值 0
	fmt.Println(ub) //初始化后的默认值 0
	fmt.Println(uc) //初始化后的默认值 0
	fmt.Println(ud) //初始化后的默认值 0

        var i int //32位的平台，为4字节，64位的平台为8字节
	var iu uint //32位的平台，为4字节，64位的平台为8字节
	fmt.Println(i) //初始化后的默认值 0
	fmt.Println(iu) //初始化后的默认值 0


//结论 int 整形初始化后默认赋值为0
}

