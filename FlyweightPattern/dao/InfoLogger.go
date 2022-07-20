package dao

import (
	"designMode/FlyweightPattern/Interface"
	"designMode/FlyweightPattern/static"
	"fmt"
)

//InfoLogger 提示日志类
type InfoLogger struct {
	Level      int
	NextLogger Interface.BaseLogger
}

//NewInfoLogger 实例化提示日志类
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      static.InfoLogLevel,
		NextLogger: nil,
	}
}

//Write 提示日志类的写方法
func (infoL *InfoLogger) Write(message string) {
	fmt.Printf("信息日志输出: %s.\n", message)
}

//PrintLog 提示日志类的输入日志方法
func (infoL *InfoLogger) PrintLog(level int, message string) {
	if infoL.Level == level {
		infoL.Write(message)
	}
	if infoL.NextLogger != nil {
		infoL.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 提示日志类设置下一个对象
func (infoL *InfoLogger) SetNextLogger(logger Interface.BaseLogger) {
	infoL.NextLogger = logger
}
