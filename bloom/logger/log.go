package logger

import "fmt"

// 日志输出
type Logger interface {
	Errorf(s string, args ...interface{})
}

// 默认日志输出
type logger struct {
}

func (l logger) Errorf(s string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(s, args...))
}

var Log Logger = &logger{}
