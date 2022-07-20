package dao

import "fmt"

type Basketball struct {
}

func NewBasketball() *Basketball {
	return &Basketball{}
}

func (*Basketball) Run() {
	fmt.Println("打篮球")
}
