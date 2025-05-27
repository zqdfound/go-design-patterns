package factory

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("This is an apple.")
}

type Peach struct{}

func (p *Peach) Show() {
	fmt.Println("This is an peach.")
}

type FruitFactory struct{}

func (f *FruitFactory) CreateFruit(fruitType string) Fruit {
	switch fruitType {
	case "apple":
		return &Apple{}
	case "peach":
		return &Peach{}
	default:
		return nil
	}
}

func ShowFruit(fruitType string) {
	factory := &FruitFactory{}
	fruit := factory.CreateFruit(fruitType)
	if fruit != nil {
		fruit.Show()
	} else {
		fmt.Println("Unknown fruit type.")
	}
}
