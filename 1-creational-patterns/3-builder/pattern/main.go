package main

import "fmt"

// HouseBuilder defines the steps to construct a house
type HouseBuilder interface {
	BuildWalls()
	BuildDoors()
	BuildWindows()
	BuildRoof()
	GetHouse() House
}

// House is the complex object being constructed
type House struct {
	Walls   string
	Doors   string
	Windows string
	Roof    string
}

// WoodenHouseBuilder builds a house with wooden materials
type WoodenHouseBuilder struct {
	house House
}

func (b *WoodenHouseBuilder) BuildWalls() {
	b.house.Walls = "Wooden Walls"
}

func (b *WoodenHouseBuilder) BuildDoors() {
	b.house.Doors = "Wooden Doors"
}

func (b *WoodenHouseBuilder) BuildWindows() {
	b.house.Windows = "Wooden Windows"
}

func (b *WoodenHouseBuilder) BuildRoof() {
	b.house.Roof = "Wooden Roof"
}

func (b *WoodenHouseBuilder) GetHouse() House {
	return b.house
}

// StoneHouseBuilder builds a house with stone materials
type StoneHouseBuilder struct {
	house House
}

func (b *StoneHouseBuilder) BuildWalls() {
	b.house.Walls = "Stone Walls"
}

func (b *StoneHouseBuilder) BuildDoors() {
	b.house.Doors = "Stone Doors"
}

func (b *StoneHouseBuilder) BuildWindows() {
	b.house.Windows = "Stone Windows"
}

func (b *StoneHouseBuilder) BuildRoof() {
	b.house.Roof = "Stone Roof"
}

func (b *StoneHouseBuilder) GetHouse() House {
	return b.house
}

// Director defines the construction process
type Director struct {
	builder HouseBuilder
}

func (d *Director) SetBuilder(builder HouseBuilder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildWalls()
	d.builder.BuildDoors()
	d.builder.BuildWindows()
	d.builder.BuildRoof()
}

func main() {
	// Create a director
	director := Director{}

	// Build a wooden house
	woodenBuilder := &WoodenHouseBuilder{}
	director.SetBuilder(woodenBuilder)
	director.Construct()
	woodenHouse := woodenBuilder.GetHouse()
	fmt.Println("Wooden House:", woodenHouse)

	// Build a stone house
	stoneBuilder := &StoneHouseBuilder{}
	director.SetBuilder(stoneBuilder)
	director.Construct()
	stoneHouse := stoneBuilder.GetHouse()
	fmt.Println("Stone House:", stoneHouse)
}
