package dao

import "fmt"

//GreenCircle 绿色圆的实体类，桥接模式接口
type GreenCircle struct{}

//NewGreenCircle 实例化绿色圆
func NewGreenCircle() *GreenCircle {
	return &GreenCircle{}
}

//DrawCircle 绿色圆实现DrawAPI方法
func (gc *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}
