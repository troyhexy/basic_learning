package object

import (
	"fmt"
	"io"
	"strings"
)

//支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段
type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S1 struct {
	*MyInt
	t
	io.Reader //Go 语言规定如果结构体使用从其他包导入的类型作为嵌入字段 比如 pkg.T，那么这个嵌入字段的字段名就是 T，代表的类型为 pkg.T
	s         string
	n         int
}

func GetEmbeddedField() {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S1{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	var sl = make([]byte, len("hello, go"))
	s.Reader.Read(sl)
	fmt.Println(string(sl)) // hello, go
	s.MyInt.Add(5)
	fmt.Println(*(s.MyInt)) // 22
}

type I interface {
	M1()
	M2()
}

type T4 struct {
	I
}

func (T4) M3() {}

//结构体类型的方法集合，包含嵌入的接口类型的方法集合。
//接口类型只能嵌入接口类型。而结构体类型对嵌入类型的要求就比较宽泛了，可以是任意自定义类型或接口类型
func GetEmbeddedMethod() {
	var t T4
	var p *T4
	dumpMethodSet(t)
	dumpMethodSet(p)
}

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type T5 struct {
	E1
	E2
}

func (T5) M1() { println("T's M1") }
func (T5) M2() { println("T's M2") }

//E1 和 E2 方法集合存在交集的情况
//为 T5 增加 M1 和 M2 方法的实现
func GetEmbeddedMethodV1() {
	t := T5{}
	t.M1()
	t.M2()
}

type T6 int
type t2 struct {
	n int
	m int
}
func (t2) Me1() { println("t2's Me1") }
type I1 interface {
	M1()
}

type S3 struct {
	T6
	*t2
	I1
	a int
	b string
}

type S2 struct {
	T6 T6
	t2 *t2
	I  I1
	a  int
	b  string
}

func GetEmbeddedMethodV2() {
	var s3 S3
	var s2 S2
	dumpMethodSet( s3 )
	dumpMethodSet( s2) //S2 不是嵌入的，相当的正常书写的字段，所有没有 对应接口的方法集合

}