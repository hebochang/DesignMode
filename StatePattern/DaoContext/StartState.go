package DaoContext

import "fmt"

//StartState 开始状态类
type StartState struct{}

//NewStartState 实例化开始状态类
func NewStartState() *StartState {
	return &StartState{}
}

//DoAction 开始状态类的DoAction，实现State接口
func (start *StartState) DoAction(context *Context) {
	fmt.Println("现在是开始状态")
	context.state = start
}

//ToString 返回开始状态类名称
func (start *StartState) ToString() string {
	return "开始状态"
}
