package Builder

import "fmt"

type SuperComputer struct {
	Name string
}

func (this *SuperComputer) buildDisk() {
	fmt.Println("超大硬盘")
}

func (this *SuperComputer) buildCPU() {
	fmt.Println("超快CPU")
}

func (this *SuperComputer) buildRom() {
	fmt.Println("超大内存")
}
