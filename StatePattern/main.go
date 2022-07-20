package main

import (
	"designMode/StatePattern/DaoContext"
	"fmt"
)

func main() {
	context := DaoContext.NewContext()

	startState := DaoContext.NewStartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())
	fmt.Println("-------------")
	stopState := DaoContext.NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
