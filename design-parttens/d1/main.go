package main

import (
	fa "design-parttens/d1/factory"
)

func main() {
	var fruitFactory = &fa.FruitFactory{}
	fruitFactory.CreateFruit("apple").Show()
	fruitFactory.CreateFruit("peach").Show()

	fa.ShowFruit("apple")
	fa.ShowFruit("peach")
}
