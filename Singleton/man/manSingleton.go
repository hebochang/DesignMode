package man

import (
	"fmt"
	"sync"
)

type man struct {
}

func (s *man) doSomething() {
	fmt.Println("do some thing")
}

var (
	once        sync.Once
	manInstance *man
)

func GetManInstance() manSingleton {
	once.Do(
		func() {
			manInstance = &man{}
		},
	)
	return manInstance
}
