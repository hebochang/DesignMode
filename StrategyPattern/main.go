package main

import (
	"designMode/StrategyPattern/dao"
	"fmt"
)

func main() {
	b := dao.Book{100}
	d65 := dao.Discount65{}
	p := b.GetPrice(&d65)
	fmt.Printf("p: %v\n", p)
	fmt.Println("-------")
	d85 := dao.Discount85{}
	p = b.GetPrice(&d85)
	fmt.Printf("p: %v\n", p)
}
