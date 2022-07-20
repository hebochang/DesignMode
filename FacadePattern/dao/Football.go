package dao

import "fmt"

type Football struct {
}

func NewFootball() *Football {
	return &Football{}
}

func (*Football) Run() {
	fmt.Println("踢足球")
}
