package single

import (
	"sync"
)

type singelton2 struct {
}

var once sync.Once

var instant2 *singelton2

func GetSingeltoneOnce() *singelton2 {
	once.Do(func() {
		instant2 = &singelton2{}
	})
	return instant2
}
