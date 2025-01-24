package main

import "fmt"

// Step 2: Tạo Subscriber Interface với phương thức update()
// observer.go: Observer
type Observer interface {
	update(string) // Phương thức được gọi khi có thông báo
	getID() string // Trả về ID của Observer
}

// Step 3: Tạo Publisher Interface với phương thức để thêm/bỏ Subscribers
// subject.go: Subject
type Subject interface {
	register(observer Observer)   // Thêm Observer vào danh sách
	deregister(observer Observer) // Xóa Observer khỏi danh sách
	notifyAll()                   // Gửi thông báo đến tất cả Observer
}

// Step 4.1: Triển khai lớp Publisher
// Concrete Subject: Item
// item.go: Concrete subject
type Item struct {
	observerList []Observer // Danh sách các Observer
	name         string     // Tên sản phẩm
	inStock      bool       // Trạng thái tồn kho
}

// Constructor cho Item
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// Cập nhật trạng thái sản phẩm và thông báo
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

// Thêm Observer vào danh sách
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

// Xóa Observer khỏi danh sách
func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

// Gửi thông báo đến tất cả Observer
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// Helper function để xóa Observer khỏi danh sách
func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Step 4.2: Triển khai lớp Subscriber
// Concrete Observer: Customer
// customer.go: Concrete observer
type Customer struct {
	id string // ID của khách hàng
}

// Nhận thông báo từ Publisher
func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

// Trả về ID của khách hàng
func (c *Customer) getID() string {
	return c.id
}

// Step 5: Tích hợp cơ chế thông báo vào Client Code
// main.go: Client code
func main() {
	// Tạo một sản phẩm mới
	shirtItem := newItem("Nike Shirt")

	// Tạo các khách hàng (Subscribers)
	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	// Đăng ký khách hàng với sản phẩm
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	// Cập nhật trạng thái sản phẩm (thông báo sẽ được gửi)
	shirtItem.updateAvailability()
}

/*
Tóm tắt các bước triển khai Observer Pattern:
1. Phân tách logic kinh doanh thành:
   - Core functionality: Publisher (Item).
   - Other logic: Subscriber (Customer).

2. Tạo Subscriber Interface với phương thức update().

3. Tạo Publisher Interface với các phương thức thêm/bỏ Subscribers.

4. Triển khai lớp Publisher và Subscriber:
   - Publisher: Item quản lý danh sách các Observer và thông báo sự kiện.
   - Subscriber: Customer nhận thông báo từ Publisher.

5. Tích hợp cơ chế thông báo vào Client Code:
   - Tạo các Subscribers (Customer) và liên kết với Publisher (Item).
   - Khi Publisher cập nhật trạng thái, tất cả Subscribers được thông báo.

Ưu điểm:
- Giảm sự phụ thuộc giữa Publisher và Subscriber.
- Dễ dàng mở rộng thêm Subscriber hoặc Publisher mới.

Nhược điểm:
- Quản lý số lượng lớn Subscribers có thể phức tạp.
- Thứ tự thông báo đến Subscribers không được đảm bảo.
*/
