package dao

import (
	"designMode/DecoratorPattern/Interface"
	"fmt"
)

//RedShapeDecorator 红色装饰器
type RedShapeDecorator struct {
	DecoratedShape Interface.Shape
}

//NewRedShapeDecorator 实例化红色装饰器
func NewRedShapeDecorator(decShape Interface.Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		DecoratedShape: decShape,
	}
}

//SetRedBorder 装饰器方法
func (rs *RedShapeDecorator) SetRedBorder() {
	fmt.Println("红色边框")
}

//Draw 实现Shape接口的方法
func (rs *RedShapeDecorator) Draw2() {
	rs.DecoratedShape.Draw2()
	rs.SetRedBorder()
}
