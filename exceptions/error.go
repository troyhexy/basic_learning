package exceptions

import (
	"errors"
	"fmt"
)

//在标准库中提供了两种方便 Go 开发者构造错误值的方法
func GetError() {
	err := errors.New("your first demo error")
	errWithCtx := fmt.Errorf("index %d is out of bounds", 1)
	println(err, errWithCtx)
}

//从 Go 1.13 版本开始，标准库 errors 包提供了 Is 函数用于错误处理方对错误值的检视

var ErrSentinel = errors.New("the underlying sentinel error")

func GetErrorIs() {
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == ErrSentinel) //false
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}

	println("err2 is not ErrSentinel")
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

//As函数类似于通过类型断言判断一个 error 类型变量是否为特定的自定义错误类型
//errors.As函数沿着 err2 所在错误链向下找到了被包装到最深处的错误值，
//并将 err2 与其类型 * MyError成功匹配。匹配成功后，errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中，
//这也是为什么println(e == err)输出 true 的原因
func GetErrorAs() {
	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)

	var e *MyError
	println(e == err)
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}
	println("MyError is not on the chain of err2")
}
