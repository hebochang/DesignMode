package dao

import (
	"designMode/FlyweightPattern/Interface"
	"designMode/FlyweightPattern/static"
	"fmt"
)

//ErrorLogger 错误日志类
type ErrorLogger struct {
	Level      int
	NextLogger Interface.BaseLogger
}

//NewErrorLogger 实例化错误日志类
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      static.ErrorLogLevel,
		NextLogger: nil,
	}
}

//Write 错误日志类写方法
func (el *ErrorLogger) Write(message string) {
	fmt.Printf("错误日志输出: %s.\n", message)
}

//PrintLog 错误日志类输入日志方法
func (el *ErrorLogger) PrintLog(level int, message string) {
	if el.Level == level {
		el.Write(message)
	}
	if el.NextLogger != nil {
		el.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 错误日志类设置下一个对象
func (el *ErrorLogger) SetNextLogger(logger Interface.BaseLogger) {
	el.NextLogger = logger
}
