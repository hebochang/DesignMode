package FactoryMode

import "fmt"

type Fruit interface {
	Grant()
	Pick()
}

type Apple struct {
}

func (receiver *Apple) Grant() {
	fmt.Println("种植苹果")
}

func (receiver *Apple) Pick() {
	fmt.Println("采摘苹果")
}

type Orange struct {
}

func (receiver *Orange) Grant() {
	fmt.Println("种植橘子")
}

func (receiver *Orange) Pick() {
	fmt.Println("采摘橘子")
}

func NewFactoryFruit(recv string) (f Fruit) {
	switch recv {
	case "apple":
		return &Apple{}
	case "orange":
		return &Orange{}
	}
	return nil
}
