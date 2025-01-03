// iSportsFactory.go: Abstract factory interface
package main

import "fmt"

// Step 3: Declare the abstract factory interface
// - Khai báo abstract factory interface để tạo tất cả các loại sản phẩm.
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// Step 5: Create factory initialization code
// - Tạo mã khởi tạo factory để khởi tạo concrete factory class dựa trên cấu hình hoặc môi trường.
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

// adidas.go: Concrete factory
// Step 4: Implement a set of concrete factory classes
// - Tạo các concrete factory class, mỗi class tương ứng với một biến thể sản phẩm.
type Adidas struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

// nike.go: Concrete factory
type Nike struct{}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

// iShoe.go: Abstract product
// Step 2: Declare abstract product interfaces
// - Khai báo các abstract product interfaces cho từng loại sản phẩm.
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Abstract product implementation
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

// adidasShoe.go: Concrete product
// Step 1: Map out a matrix of distinct product types versus variants of these products
// - AdidasShoe represents a specific product variant.
type AdidasShoe struct {
	Shoe
}

// nikeShoe.go: Concrete product
// - NikeShoe represents a specific product variant.
type NikeShoe struct {
	Shoe
}

// iShirt.go: Abstract product
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Abstract product implementation
type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

// adidasShirt.go: Concrete product
type AdidasShirt struct {
	Shirt
}

// nikeShirt.go: Concrete product
type NikeShirt struct {
	Shirt
}

// main.go: Client code
func main() {
	// Step 5: Create factory initialization code
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	// Step 6: Replace direct product constructors with factory methods
	// - Sử dụng factory methods thay vì gọi trực tiếp product constructors.
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	// Print details for products
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

// Helper functions to print product details
func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
