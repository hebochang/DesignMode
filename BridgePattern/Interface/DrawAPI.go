package Interface

//DrawAPI 画图抽象接口，桥接模式的抽象接口
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}
