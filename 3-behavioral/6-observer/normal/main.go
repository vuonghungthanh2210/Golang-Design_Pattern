// observer_without_pattern.go: Implementation without Observer Pattern

package main

import "fmt"

// Item quản lý danh sách khách hàng trực tiếp
type Item struct {
	name      string
	inStock   bool
	customers []*Customer // Danh sách khách hàng
}

// Customer lưu thông tin của từng khách hàng
type Customer struct {
	id string
}

// Constructor cho Item
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// Constructor cho Customer
func newCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

// Thêm khách hàng trực tiếp vào danh sách
func (i *Item) addCustomer(customer *Customer) {
	i.customers = append(i.customers, customer)
}

// Xóa khách hàng khỏi danh sách
func (i *Item) removeCustomer(customer *Customer) {
	for index, c := range i.customers {
		if c.id == customer.id {
			i.customers = append(i.customers[:index], i.customers[index+1:]...)
			break
		}
	}
}

// Cập nhật trạng thái sản phẩm và thông báo khách hàng trực tiếp
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	for _, customer := range i.customers {
		fmt.Printf("Sending email to customer %s for item %s\n", customer.id, i.name)
	}
}

func main() {
	// Tạo sản phẩm và khách hàng
	shirtItem := newItem("Nike Shirt")
	customer1 := newCustomer("abc@gmail.com")
	customer2 := newCustomer("xyz@gmail.com")

	// Thêm khách hàng vào danh sách sản phẩm
	shirtItem.addCustomer(customer1)
	shirtItem.addCustomer(customer2)

	// Cập nhật trạng thái sản phẩm
	shirtItem.updateAvailability()
}

/*
Nhược điểm và khó khăn khi không dùng Design Pattern:

1. **Phụ thuộc chặt chẽ giữa các lớp**:
   - `Item` phụ thuộc trực tiếp vào danh sách `Customer`.
   - Thay đổi logic thông báo sẽ ảnh hưởng đến `Item`.

2. **Khó mở rộng**:
   - Nếu muốn thêm logic mới (ví dụ: gửi thông báo qua SMS hoặc push notification), cần thay đổi trực tiếp vào lớp `Item`.

3. **Khó tái sử dụng**:
   - Lớp `Customer` bị gắn chặt với `Item`. Không thể tái sử dụng `Customer` trong ngữ cảnh khác mà không kéo theo logic liên quan.

4. **Quản lý phức tạp**:
   - Danh sách khách hàng được quản lý trực tiếp trong `Item`. Khi số lượng lớn, việc xử lý thêm/xóa hoặc thông báo sẽ phức tạp và dễ lỗi.

5. **Vi phạm nguyên tắc Single Responsibility Principle**:
   - Lớp `Item` đảm nhận nhiều nhiệm vụ: quản lý trạng thái sản phẩm, quản lý danh sách khách hàng và gửi thông báo.

Khi sử dụng Observer Pattern, những vấn đề này được giải quyết bằng cách phân tách trách nhiệm giữa `Publisher` và `Subscriber`. Các thành phần này được tách biệt, linh hoạt và dễ mở rộng hơn.
*/
