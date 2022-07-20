package Interface

import (
	"designMode/StatePattern/DaoContext"
)

//State 状态接口
type State interface {
	DoAction(context *DaoContext.Context)
	ToString() string
}
