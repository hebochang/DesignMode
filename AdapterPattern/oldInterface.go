package main

import "fmt"

type OldInterface interface {
	OldMethod()
}

type OldImpl struct {
}

func (OldImpl) OldMethod() {
	fmt.Println("旧方法实现")
}
