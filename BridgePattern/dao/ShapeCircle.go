package dao

import "designMode/BridgePattern/Interface"

//ShapeCircle 桥接模式的实体类
type ShapeCircle struct {
	Radius  int
	X       int
	Y       int
	drawAPI Interface.DrawAPI
}

//NewShapeCircle 实例化桥接模式实体类
func NewShapeCircle(radius, x, y int, drawAPI Interface.DrawAPI) *ShapeCircle {
	return &ShapeCircle{
		Radius:  radius,
		X:       x,
		Y:       y,
		drawAPI: drawAPI,
	}
}

//Draw 实体类的Draw方法
func (sc *ShapeCircle) Draw1() {
	sc.drawAPI.DrawCircle(sc.Radius, sc.X, sc.Y)
}
