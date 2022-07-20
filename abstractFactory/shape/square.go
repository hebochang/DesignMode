package shape

import "fmt"

type Square struct{}

func (*Square) Draw() {
	fmt.Println("Draw Square")
}
