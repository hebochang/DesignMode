package dao

import "fmt"

//RedCircle 红色圆的类，桥接模式接口
type RedCircle struct{}

//NewRedCircle 实例化红色圆
func NewRedCircle() *RedCircle {
	return &RedCircle{}
}

//DrawCircle 红色圆实现DrawAPI方法
func (rc *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}
