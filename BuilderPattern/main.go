package main

import (
	"designMode/BuilderPattern/Builder"
	"fmt"
)

func main() {
	sc := Builder.SuperComputer{}
	d := Builder.NewConstruct(&sc)
	d.Consturct()

	fmt.Println("--------------")

	lc := Builder.LowComputer{}
	d2 := Builder.NewConstruct(&lc)
	d2.Consturct()
}
