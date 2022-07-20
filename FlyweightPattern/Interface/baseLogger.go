package Interface

//BaseLogger 日志接口
type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}
