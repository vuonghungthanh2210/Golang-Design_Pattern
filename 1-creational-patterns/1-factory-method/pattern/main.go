package main

import (
	"fmt"
)

// Step 1: Make all products follow the same interface
// Define a common interface for all products (IGun).
// iGun.go: Product interface
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Step 1: Concrete product implementing the interface
// Base struct Gun that implements the common methods of IGun.
// gun.go: Concrete product
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Step 4: Create a set of creator subclasses
// Concrete product Ak47 implementing the IGun interface.
// ak47.go: Concrete product
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Concrete product Musket implementing the IGun interface.
// musket.go: Concrete product
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// Step 2: Add an empty factory method inside the creator class
// Factory method to return IGun interface objects.
// gunFactory.go: Factory
func getGun(gunType string) (IGun, error) {
	// Step 5: Reuse the control parameter to return the right product
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

// Step 3: Replace constructors with calls to the factory method
// main.go: Client code
func main() {
	// Client code using the factory method to create objects
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	// Display the details of each product
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

// Ví Dụ Mở Rộng
// 	Nếu sau này cần thêm loại súng mới, chẳng hạn như Sniper:
// 		Tạo class mới Sniper kế thừa giao diện chung.
// 		Cập nhật logic trong getGun để trả về Sniper khi nhận được tham số "sniper".
// 		Không cần thay đổi bất kỳ đoạn mã nào trong mã client.
// Đảm bảo tính tái sử dụng:
// 	Sử dụng Factory Method để quản lý mọi loại súng, giúp đơn giản hóa quá trình mở rộng và bảo trì mã nguồn.
