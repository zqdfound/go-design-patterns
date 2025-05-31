package action

import "fmt"

//observe design pattern
//Simulated stock trading
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyAll()
}
type Observer interface {
	Update(price float64) bool
	GetId() int
}

// Concrete Observer
type Stock struct {
	observers []Observer
	price     float64
	name      string
}

func NewStock(name string) *Stock {
	return &Stock{
		observers: make([]Observer, 0),
		name:      name,
	}
}
func (s *Stock) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}
func (s *Stock) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs.GetId() == observer.GetId() {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}
func (s *Stock) NotifyAll() {
	removeList := make([]Observer, 0)
	for _, observer := range s.observers {
		if observer.Update(s.price) {
			removeList = append(removeList, observer)
		}
	}
	// Remove observers that returned true from Update
	for _, observer := range removeList {
		s.RemoveObserver(observer)
	}
}

func (s *Stock) UpdatePrice(price float64) {
	s.price = price
	fmt.Printf("Stock %s price updated to %.2f\n", s.name, price)
	s.NotifyAll()
}

//Investor is an Observer
type Investor struct {
	id        int
	sellPrice float64
}

func NewInvestor(id int, price float64) *Investor {
	return &Investor{
		id:        id,
		sellPrice: price,
	}
}
func (i *Investor) Update(price float64) bool {
	fmt.Printf("Investor %d: Stock price updated to %.2f\n", i.id, price)
	if price >= i.sellPrice {
		fmt.Printf("Investor %d: Selling stock at price %.2f\n", i.id, price)
		return true // Return true to indicate that the observer should be removed
	} else {
		fmt.Printf("Investor %d: Holding stock, current price %.2f is below sell price %.2f\n", i.id, price, i.sellPrice)
		return false
	}
}
func (i *Investor) GetId() int {
	return i.id
}
