package main

import "fmt"

// House represents the final product
type House struct {
	Walls   string
	Doors   string
	Windows string
	Roof    string
}

func NewHouse(walls, doors, windows, roof string) House {
	return House{
		Walls:   walls,
		Doors:   doors,
		Windows: windows,
		Roof:    roof,
	}
}

func main() {
	// Build a Wooden House
	woodenHouse := House{
		Walls:   "Wooden Walls",
		Doors:   "Wooden Doors",
		Windows: "Wooden Windows",
		Roof:    "Wooden Roof",
	}
	fmt.Println("Wooden House:", woodenHouse)

	// Build a Stone House
	stoneHouse := House{
		Walls:   "Stone Walls",
		Doors:   "Stone Doors",
		Windows: "Stone Windows",
		Roof:    "Stone Roof",
	}
	fmt.Println("Stone House:", stoneHouse)

	// Build a Custom House
	customHouse := NewHouse("Brick Walls", "Glass Doors", "Large Windows", "Green Roof")
	fmt.Println("Custom House:", customHouse)
}
