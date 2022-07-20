package main

import "fmt"

//适配器struct集成oldInterface，并且实现NewInterface
type Adapter struct {
	OldInterface
}

// 适配器
func (a *Adapter) NewMethod() {
	fmt.Println("新方法实现")
	a.OldMethod()
}

func main() {
	oldInteface := OldImpl{}
	a := Adapter{OldInterface: oldInteface}
	a.NewMethod()
	a.OldMethod()
}
