package object

import (
	"fmt"
	"reflect"
)

//    • 类型 T 方法集包含全部 receiver T 方法。
//    • 类型 *T 方法集包含全部 receiver T + *T 方法。
//    • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
//    • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
//    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。




type T3 struct{}

func (T3) M1() {}
func (T3) M2() {}

func (*T3) M3() {}
func (*T3) M4() {}

func GetMethodCollect() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T3
	dumpMethodSet(t)
	dumpMethodSet(&t)
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}


