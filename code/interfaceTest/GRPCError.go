package interfaceTest

import "fmt"

// GRPCError 结构体定义了一个具体的错误类型，实现了 error 接口
type GRPCError struct {
	Code    int    // 错误代码
	Message string // 错误消息
}

// Error 方法实现了 error 接口的 Error 方法，用于返回错误消息
func (e *GRPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

// NewError 函数用于创建一个 GRPCError 实例，并将其作为 error 接口类型返回
func NewError(code int, message string) error {
	return &GRPCError{
		Code:    code,
		Message: message,
	}
}

// AsError 函数用于将一个 error 接口类型的值直接返回，不进行任何转换
func AsError(err error) error {
	return err
}

// PrintInterface 函数接受一个空接口 v 作为参数，并将其内容打印出来。
// 由于空接口可以接受任何类型的值，因此这个函数可以打印任意类型的值。类似于其他语言中的泛型函数。
func PrintInterface(v interface{}) {
	fmt.Println(v)
}
