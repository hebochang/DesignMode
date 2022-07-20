package color

import "fmt"

type Green struct{}

func (*Green) Fill() {
	fmt.Println("Fill Green")
}
