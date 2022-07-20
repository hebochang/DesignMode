package main

import "fmt"

func main() {
	cpu := CPU{"奔腾586"}
	rom := ROM{"金士顿"}
	disk := Disk{"三星"}

	c := Computer{
		Cpu:  cpu,
		Rom:  rom,
		Disk: disk,
	}

	c1 := c.BackUp()
	fmt.Printf("c1: %v\n", *c1)

}
