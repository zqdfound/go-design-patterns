package single

import "fmt"

type singelton struct {
}

// var instant *singelton = &singelton{}
var instant *singelton

// func GetSingeltone() *singelton {
// 	return instant
// }

func GetSingeltoneDelay() *singelton {
	if instant == nil {
		instant = &singelton{}
	}
	return instant
}

func (s *singelton) DoSomething() {
	fmt.Println("Doing something in singleton")
}
