package exceptions

import (
	"fmt"
)

//panic 函数支持的参数类型是 interface{}
func GetPanic() {


	panic("除数不能为0！")
	//panic(errors.New("除数不能为0")) // 传入 error 类型
}

//recover 是 Go 内置的专门用于恢复 panic 的函数，
//它必须被放在一个 defer 函数中才能生效。
//如果 recover 捕捉到 panic，它就会返回以 panic 的具体内容为错误上下文信息的错误值。
//如果没有 panic 发生，那么 recover 将返回 nil。
//而且，如果 panic 被 recover 捕捉到，panic 引发的 panicking 过程就会停止。

func GetRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()

	panic("这是个错误")

}
