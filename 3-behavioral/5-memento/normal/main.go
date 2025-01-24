package main

import "fmt"

// Document lưu trạng thái hiện tại và lịch sử của tài liệu
type Document struct {
	stateHistory []string // Lưu lịch sử các trạng thái
	state        string   // Trạng thái hiện tại
}

// Thay đổi trạng thái của tài liệu và lưu trạng thái cũ vào lịch sử
func (d *Document) setState(newState string) {
	d.stateHistory = append(d.stateHistory, d.state) // Lưu trạng thái cũ
	d.state = newState                               // Cập nhật trạng thái mới
}

// Quay lại trạng thái trước đó
func (d *Document) undo() {
	if len(d.stateHistory) == 0 {
		fmt.Println("Không có trạng thái trước đó để hoàn tác.")
		return
	}
	d.state = d.stateHistory[len(d.stateHistory)-1]         // Lấy trạng thái trước đó
	d.stateHistory = d.stateHistory[:len(d.stateHistory)-1] // Xóa trạng thái cuối cùng khỏi lịch sử
}

// Lấy trạng thái hiện tại
func (d *Document) getState() string {
	return d.state
}

// Main function
func main() {
	document := &Document{
		state: "State A", // Khởi tạo với trạng thái ban đầu
	}

	fmt.Printf("Trạng thái ban đầu: %s\n", document.getState())

	// Thay đổi trạng thái và lưu lịch sử
	document.setState("State B")
	fmt.Printf("Trạng thái mới: %s\n", document.getState())

	document.setState("State C")
	fmt.Printf("Trạng thái mới: %s\n", document.getState())

	// Hoàn tác trạng thái
	document.undo()
	fmt.Printf("Trạng thái sau hoàn tác: %s\n", document.getState())

	document.undo()
	fmt.Printf("Trạng thái sau hoàn tác: %s\n", document.getState())

	// Thử hoàn tác khi không còn trạng thái
	document.undo()
}

// Nhược điểm và Khó khăn của cách tiếp cận này (Không dùng Design Pattern):

/*
1. Phá vỡ tính đóng gói (Encapsulation):
   - Toàn bộ logic lưu trữ và quản lý lịch sử trạng thái đều nằm trong lớp `Document`.
   - Điều này làm lớp này phình to và khó bảo trì.

2. Khó tái sử dụng:
   - Logic quản lý lịch sử không được tách biệt, nên không thể tái sử dụng trong các lớp khác
     có nhu cầu lưu trạng thái tương tự.

3. Thiếu an toàn dữ liệu:
   - Lịch sử trạng thái (`stateHistory`) là một mảng được quản lý trực tiếp trong lớp `Document`.
   - Dễ dẫn đến chỉnh sửa sai hoặc mất dữ liệu do các thành phần khác trong ứng dụng can thiệp.

4. Khó mở rộng:
   - Khi cần thêm tính năng mới (như redo), phải sửa đổi trực tiếp lớp `Document`.
   - Dẫn đến rủi ro phá vỡ các tính năng hiện có.

5. Khó bảo trì:
   - Tất cả các logic liên quan đến lưu trữ, khôi phục và chỉnh sửa trạng thái đều tập trung
     trong một lớp duy nhất. Điều này làm tăng độ phức tạp và giảm khả năng đọc hiểu mã nguồn.

6. Không tuân thủ nguyên tắc Single Responsibility Principle (SRP):
   - Lớp `Document` đảm nhận cả việc quản lý trạng thái hiện tại và lịch sử, làm sai lệch mục đích chính
     của nó là đại diện cho tài liệu.

7. Không có cơ chế kiểm soát lịch sử:
   - Lớp `Document` không thể đặt giới hạn về số lượng trạng thái được lưu hoặc xóa lịch sử cũ
     một cách tự động, có thể gây lãng phí bộ nhớ.

8. Tăng độ phức tạp khi cần tích hợp thêm chức năng:
   - Nếu cần tích hợp thêm các chức năng khác như lưu lịch sử dưới dạng file hoặc đồng bộ hóa trạng thái,
     lớp `Document` sẽ phải sửa đổi lớn, vi phạm Open/Closed Principle.
*/

// Kết luận:
// Việc không áp dụng Design Pattern (Memento) dẫn đến mã nguồn phức tạp hơn, khó bảo trì và khó mở rộng.
// Sử dụng Memento Pattern sẽ giúp tách biệt rõ ràng các trách nhiệm, tăng tính bảo trì, mở rộng và an toàn dữ liệu.
