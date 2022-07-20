package color

import "fmt"

type Red struct{}

func (*Red) Fill() {
	fmt.Println("Fill Red")
}
