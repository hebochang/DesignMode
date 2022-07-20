package dao

import (
	"designMode/FlyweightPattern/Interface"
	"designMode/FlyweightPattern/static"
	"fmt"
)

//StandardLogger 标准日志类
type StandardLogger struct {
	Level      int
	NextLogger Interface.BaseLogger
}

//NewStandardLogger 实例化标准日志类
func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		Level:      static.StandardLogLevel,
		NextLogger: nil,
	}
}

//Write 标准日志类写日志方法
func (sl *StandardLogger) Write(message string) {
	fmt.Printf("标准日志输出: %s.\n", message)
}

//PrintLog 标准日志类输入日志，并且流向下一个对象方法
func (sl *StandardLogger) PrintLog(level int, message string) {
	if sl.Level == level {
		sl.Write(message)
	}
	if sl.NextLogger != nil {
		sl.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 标准日志类设置下一个对象方法
func (sl *StandardLogger) SetNextLogger(logger Interface.BaseLogger) {
	sl.NextLogger = logger
}
