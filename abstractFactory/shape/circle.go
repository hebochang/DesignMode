package shape

import "fmt"

type Circle struct{}

func (*Circle) Draw() {
	fmt.Println("Draw Circle")
}
