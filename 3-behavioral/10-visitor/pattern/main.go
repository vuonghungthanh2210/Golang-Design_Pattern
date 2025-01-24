// Bước 1: Tạo Visitor Interface với các phương thức cho từng loại đối tượng cụ thể.
package main

import "fmt"

// Visitor định nghĩa các phương thức cho từng loại đối tượng cụ thể.
// visitor.go: Visitor
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// Bước 2: Cập nhật Element Interface để thêm phương thức accept() nhận Visitor.
// Element Interface đại diện cho các đối tượng có thể chấp nhận Visitor.
// shape.go: Element
type Shape interface {
	getType() string
	accept(Visitor)
}

// Bước 3: Triển khai phương thức accept() trong từng lớp cụ thể và gọi phương thức tương ứng trong Visitor.
// Square đại diện cho hình vuông.
// square.go: Concrete element
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s) // Gọi phương thức phù hợp trong Visitor.
}

func (s *Square) getType() string {
	return "Square"
}

// Circle đại diện cho hình tròn.
// circle.go: Concrete element
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c) // Gọi phương thức phù hợp trong Visitor.
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle đại diện cho hình chữ nhật.
// rectangle.go: Concrete element
type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r) // Gọi phương thức phù hợp trong Visitor.
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

// Bước 4: Tạo các lớp Visitor cụ thể để thực hiện hành vi mong muốn.
// AreaCalculator tính diện tích của các đối tượng.
// areaCalculator.go: Concrete visitor
type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Tính diện tích hình vuông")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Tính diện tích hình tròn")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Tính diện tích hình chữ nhật")
}

// MiddleCoordinates tính tọa độ trung tâm của các đối tượng.
// middleCoordinates.go: Concrete visitor
type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Tính tọa độ trung tâm hình vuông")
}

func (m *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Tính tọa độ trung tâm hình tròn")
}

func (m *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("Tính tọa độ trung tâm hình chữ nhật")
}

// Bước 5: Sử dụng client để duyệt qua các đối tượng và gọi accept() với Visitor.
// main.go: Client code
func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)    // Tính diện tích hình vuông.
	circle.accept(areaCalculator)    // Tính diện tích hình tròn.
	rectangle.accept(areaCalculator) // Tính diện tích hình chữ nhật.

	fmt.Println()

	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)    // Tính tọa độ trung tâm hình vuông.
	circle.accept(middleCoordinates)    // Tính tọa độ trung tâm hình tròn.
	rectangle.accept(middleCoordinates) // Tính tọa độ trung tâm hình chữ nhật.
}

// Tóm tắt các bước:
// 1. Tạo Visitor Interface với các phương thức cho từng loại đối tượng cụ thể.
// 2. Cập nhật Element Interface để thêm phương thức accept() nhận Visitor.
// 3. Triển khai phương thức accept() trong từng lớp cụ thể và gọi phương thức tương ứng trong Visitor.
// 4. Tạo các lớp Visitor cụ thể để thực hiện hành vi mong muốn.
// 5. Sử dụng client để duyệt qua các đối tượng và gọi accept() với Visitor.
