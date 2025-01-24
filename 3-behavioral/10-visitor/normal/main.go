package main

import "fmt"

// Nếu không sử dụng Visitor Pattern, các hành vi như tính diện tích hoặc tọa độ trung tâm sẽ được thêm trực tiếp vào các lớp hình học.

type Shape interface {
	getType() string
	getArea() float64
	getMiddleCoordinates() (float64, float64)
}

// Hình vuông

type Square struct {
	side float64
}

func (s *Square) getType() string {
	return "Square"
}

func (s *Square) getArea() float64 {
	return s.side * s.side
}

func (s *Square) getMiddleCoordinates() (float64, float64) {
	return s.side / 2, s.side / 2
}

// Hình tròn

type Circle struct {
	radius float64
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) getMiddleCoordinates() (float64, float64) {
	return 0, 0
}

// Hình chữ nhật

type Rectangle struct {
	length  float64
	breadth float64
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

func (r *Rectangle) getArea() float64 {
	return r.length * r.breadth
}

func (r *Rectangle) getMiddleCoordinates() (float64, float64) {
	return r.length / 2, r.breadth / 2
}

// Client code: Gọi các phương thức trực tiếp từ các lớp hình học

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{length: 4, breadth: 5}

	shapes := []Shape{square, circle, rectangle}

	for _, shape := range shapes {
		fmt.Printf("%s Area: %.2f\n", shape.getType(), shape.getArea())
		x, y := shape.getMiddleCoordinates()
		fmt.Printf("%s Middle Coordinates: (%.2f, %.2f)\n", shape.getType(), x, y)
	}
}

/*
Nhược điểm và khó khăn:
1. **Vi phạm nguyên lý mở/đóng (Open/Closed Principle):**
   - Khi cần thêm hành vi mới (ví dụ: tính chu vi), ta phải chỉnh sửa tất cả các lớp hình học để thêm phương thức mới, làm tăng rủi ro bug.

2. **Phụ thuộc nhiều vào các lớp cụ thể:**
   - Khi client muốn xử lý các hình theo cách riêng, chúng phải biết chi tiết từng lớp, gây khó khăn trong việc mở rộng.

3. **Code trùng lặp:**
   - Các hành vi tương tự (ví dụ: tính diện tích) có thể có logic giống nhau giữa các lớp, dẫn đến trùng lặp mã nguồn.

4. **Khó bảo trì:**
   - Khi số lượng hình học hoặc hành vi tăng lên, các lớp trở nên phức tạp và khó bảo trì hơn.

5. **Không linh hoạt:**
   - Không thể thêm hành vi mới mà không chỉnh sửa các lớp hình học hiện tại, gây ảnh hưởng đến tính ổn định của hệ thống.

Lợi ích của Visitor Pattern so với cách tiếp cận này:
- Tách biệt dữ liệu (Shape) và hành vi (Visitor).
- Dễ dàng mở rộng thêm hành vi mà không cần thay đổi các lớp hiện tại.
- Giảm phụ thuộc giữa các lớp và đảm bảo nguyên lý mở/đóng.
*/
