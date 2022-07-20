package main

import "designMode/DecoratorPattern/dao"

func main() {
	c := dao.NewCircle()
	rsd := dao.NewRedShapeDecorator(c)
	rsd.Draw2()
}
