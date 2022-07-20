package DaoContext

import "fmt"

//StopState 停止状态类
type StopState struct{}

//NewStopState 实例化停止状态类
func NewStopState() *StopState {
	return &StopState{}
}

//DoAction 停止状态类方法，实现State接口
func (stop *StopState) DoAction(context *Context) {
	fmt.Println("现在是停止状态")
	context.state = stop
}

//ToString 返回停止状态类名称
func (stop *StopState) ToString() string {
	return "停止状态"
}
