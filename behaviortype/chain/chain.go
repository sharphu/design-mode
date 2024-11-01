/*
概念与定义：责任链模式将处理请求的对象串联成一条链，并沿着链传递请求，直到某个对象处理它。
例子：使用一个支持不同级别的日志系统，包含 Info、Debug 和 Error 三个级别的日志处理器。每个处理器负责不同级别的日志输出，可以按需设置下一个处
理器。
*/

package chain

import "fmt"

// Logger 定义日志接口
type Logger interface {
	SetNext(Logger)
	Log(level int, message string)
}

// LogLevel 常量定义日志级别
const (
	InfoLevel = iota
	DebugLevel
	ErrorLevel
)

// BaseLogger 基础日志结构体，包含下一个处理器
type BaseLogger struct {
	next Logger
}

// SetNext 设置下一个处理器
func (b *BaseLogger) SetNext(next Logger) {
	b.next = next
}

// InfoLogger 处理 Info 级别日志的结构体
type InfoLogger struct {
	BaseLogger
}

// NewInfoLogger InfoLogger 的构造函数
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{}
}

// Log 处理 Info 级别日志
func (i *InfoLogger) Log(level int, message string) {
	if level == InfoLevel {
		fmt.Println("INFO:", message)
	} else if i.next != nil {
		i.next.Log(level, message)
	}
}

// DebugLogger 处理 Debug 级别日志的结构体
type DebugLogger struct {
	BaseLogger
}

// NewDebugLogger DebugLogger 的构造函数
func NewDebugLogger() *DebugLogger {
	return &DebugLogger{}
}

// Log 处理 Debug 级别日志
func (d *DebugLogger) Log(level int, message string) {
	if level == DebugLevel {
		fmt.Println("DEBUG:", message)
	} else if d.next != nil {
		d.next.Log(level, message)
	}
}

// ErrorLogger 处理 Error 级别日志的结构体
type ErrorLogger struct {
	BaseLogger
}

// NewErrorLogger ErrorLogger 的构造函数
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{}
}

// Log 处理 Error 级别日志
func (e *ErrorLogger) Log(level int, message string) {
	if level == ErrorLevel {
		fmt.Println("ERROR:", message)
	} else if e.next != nil {
		e.next.Log(level, message)
	}
}
