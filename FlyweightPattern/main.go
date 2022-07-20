package main

import (
	"designMode/FlyweightPattern/Interface"
	"designMode/FlyweightPattern/dao"
	"designMode/FlyweightPattern/static"
)

//GetChainOfLoggers 获取日志对象链
func GetChainOfLoggers() Interface.BaseLogger {
	errLog := dao.NewErrorLogger()
	infoLog := dao.NewInfoLogger()
	standardLog := dao.NewStandardLogger()

	errLog.SetNextLogger(infoLog)
	infoLog.SetNextLogger(standardLog)

	return errLog
}

func main() {
	logChain := GetChainOfLoggers()
	logChain.PrintLog(static.InfoLogLevel, "这是一条信息")
	logChain.PrintLog(static.StandardLogLevel, "这是标准输出")
	logChain.PrintLog(static.ErrorLogLevel, "这是错误信息")
}
