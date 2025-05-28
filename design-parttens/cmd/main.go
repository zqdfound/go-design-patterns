package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Println("初始化操作，只会执行一次")
}

func main() {
	// 多次调用，但 setup() 只会执行一次
	once.Do(setup)
	once.Do(setup)
	once.Do(setup)
}
