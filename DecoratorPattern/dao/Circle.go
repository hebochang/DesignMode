package dao

import "fmt"

//Circle 圆形类
type Circle struct{}

//NewCircle 实例化圆形
func NewCircle() *Circle {
	return &Circle{}
}

//Draw 输出方法，实现Shape接口
func (c *Circle) Draw2() {
	fmt.Println("画圆方法")
}
