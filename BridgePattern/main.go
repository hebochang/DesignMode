package main

import "designMode/BridgePattern/dao"

//我们使用DrawAPI作为桥接模式的抽象接口，ShapeCirlce作为桥接模式的实体类。
//将抽象接口保存在实体类中，使得抽象接口实例变化时，ShapeCircle可以始终不变。
func main() {
	redCircle := dao.NewShapeCircle(5, 6, 8, dao.NewRedCircle())
	redCircle.Draw1()

	greenCircle := dao.NewShapeCircle(1, 2, 4, dao.NewGreenCircle())
	greenCircle.Draw1()
}
