package main

import "fmt"

// Không sử dụng Iterator Pattern
// Duyệt qua danh sách người dùng trực tiếp thông qua vòng lặp

type User struct {
	name string
	age  int
}

func main() {
	// Tạo danh sách người dùng
	users := []*User{
		{name: "Alice", age: 30},
		{name: "Bob", age: 25},
		{name: "Charlie", age: 35},
	}

	// Duyệt qua danh sách người dùng
	for i := 0; i < len(users); i++ {
		fmt.Printf("User: %+v\n", users[i])
	}

	// Thay đổi cách duyệt qua danh sách nếu có yêu cầu khác
	fmt.Println("Duyệt ngược qua danh sách người dùng:")
	for i := len(users) - 1; i >= 0; i-- {
		fmt.Printf("User: %+v\n", users[i])
	}
}

/*
Nhược điểm và khó khăn khi không sử dụng Design Pattern:
1. Phụ thuộc vào cấu trúc dữ liệu cụ thể:
   - Mã client phải biết chi tiết về cách danh sách người dùng được lưu trữ (ở đây là một mảng).
   - Nếu thay đổi cấu trúc lưu trữ (ví dụ, từ mảng sang danh sách liên kết hoặc cây), mã client phải được sửa đổi.

2. Thiếu khả năng tái sử dụng:
   - Logic duyệt qua danh sách (như vòng lặp `for`) không được đóng gói trong một lớp riêng biệt. Điều này dẫn đến việc lặp lại mã nếu cần duyệt danh sách ở nhiều nơi khác nhau.

3. Khó mở rộng:
   - Nếu cần thêm các kiểu duyệt khác (như duyệt theo điều kiện hoặc duyệt song song), phải thay đổi hoặc thêm logic trong mã client, làm tăng độ phức tạp.

4. Không hỗ trợ duyệt độc lập:
   - Nếu cần nhiều đối tượng độc lập duyệt qua cùng một danh sách, trạng thái của các vòng lặp không được quản lý tách biệt, dễ dẫn đến lỗi.

5. Khó bảo trì:
   - Việc thay đổi hoặc thêm logic duyệt qua có thể ảnh hưởng đến các phần khác trong mã, làm tăng rủi ro lỗi.

6. Không đáp ứng nguyên tắc trách nhiệm đơn lẻ (Single Responsibility Principle):
   - Mã client chịu trách nhiệm cả về logic duyệt qua lẫn logic xử lý từng phần tử, dẫn đến việc không tách bạch rõ ràng.

Tóm lại, việc không sử dụng Design Pattern dẫn đến mã khó bảo trì, mở rộng và tái sử dụng, đặc biệt khi yêu cầu phát triển hoặc thay đổi cấu trúc dữ liệu.
*/
