package main

// import (
// 	fa "design-parttens/d1/factory"
// )

// func main() {
// 	var fruitFactory = &fa.FruitFactory{}
// 	fruitFactory.CreateFruit("apple").Show()
// 	fruitFactory.CreateFruit("peach").Show()

// 	fa.ShowFruit("apple")
// 	fa.ShowFruit("peach")
// }

import (
	si "design-parttens/d1/single"
	"fmt"
)

func main() {
	// s.DoSomething()
	for i := 0; i < 10; i++ {
		s := si.GetSingeltoneDelay()
		fmt.Println(&s)
	}

}
