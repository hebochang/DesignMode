package DaoContext

import "designMode/StatePattern/Interface"

//Context 状态保存类
type Context struct {
	state Interface.State
}

//NewContext 实例化状态保存类
func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

//SetState 设置状态保存类当前的状态
func (c *Context) SetState(s Interface.State) {
	c.state = s
}

//GetState 获取状态保存类当前的状态
func (c *Context) GetState() Interface.State {
	return c.state
}
