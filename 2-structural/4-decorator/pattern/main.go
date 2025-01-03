package main

import "fmt"

// Step 1: Đảm bảo rằng miền nghiệp vụ của bạn có thể được biểu diễn dưới dạng một thành phần chính với nhiều lớp tùy chọn bao quanh nó.
// Trong ví dụ này, thành phần chính là một loại pizza cơ bản, và các lớp tùy chọn là các topping như phô mai và cà chua.

type IPizza interface {
	getPrice() int
}

// Concrete Component: VeggieMania

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15 // Giá cơ bản của pizza là 15.
}

// Concrete Component: Hamburger
type Hamburger struct{}

func (h *Hamburger) getPrice() int {
	return 20 // Giá cơ bản của hamburger là 20.
}

// Base Decorator
type BaseTopping struct {
	pizza IPizza
}

func (b *BaseTopping) getPrice() int {
	return b.pizza.getPrice()
}

// Concrete Decorators: TomatoTopping

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	return c.pizza.getPrice() + 7
}

// Concrete Decorators: CheeseTopping
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}

func main() {
	// Tạo pizza cơ bản
	pizza := &VeggieMania{}

	// Thêm topping phô mai
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	// Thêm topping cà chua
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	// Tạo hamburger cơ bản
	hamburger := &Hamburger{}

	// Thêm topping phô mai cho hamburger
	hamburgerWithCheese := &CheeseTopping{
		pizza: hamburger,
	}

	// Thêm topping cà chua cho hamburger
	hamburgerWithCheeseAndTomato := &TomatoTopping{
		pizza: hamburgerWithCheese,
	}

	// Hiển thị giá pizza
	fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

	// Hiển thị giá hamburger
	fmt.Printf("Price of Hamburger with tomato and cheese topping is %d\n", hamburgerWithCheeseAndTomato.getPrice())
}
