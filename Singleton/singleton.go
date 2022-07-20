package main

import (
	"designMode/Singleton/man"
	"fmt"
)

func main() {
	s1 := man.GetManInstance()
	fmt.Printf("s1: %p\n", s1)
	s2 := man.GetManInstance()
	fmt.Printf("s2: %p\n", s2)
}
