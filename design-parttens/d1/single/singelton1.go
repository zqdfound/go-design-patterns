package single

import (
	"sync"
)

type singelton1 struct {
}

var lock sync.Mutex

var instant1 *singelton1

func GetSingeltoneLock() *singelton1 {
	lock.Lock()
	defer lock.Unlock()
	if instant1 == nil {
		instant1 = &singelton1{}
	}
	return instant1
}
