package Builder

import "fmt"

type LowComputer struct {
	Name string
}

func (this *LowComputer) buildDisk() {
	fmt.Println("超小硬盘")
}

func (this *LowComputer) buildCPU() {
	fmt.Println("超慢CPU")
}

func (this *LowComputer) buildRom() {
	fmt.Println("超小内存")
}
