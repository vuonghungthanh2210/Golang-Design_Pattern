package main

import "fmt"

// Định nghĩa giao diện IGun
type IGun interface {
	getName() string
	getPower() int
}

// Lớp chung Gun
type Gun struct {
	name  string
	power int
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// Các loại súng cụ thể
type Ak47 struct {
	Gun
}

type Musket struct {
	Gun
}

// Hàm chính để tạo đối tượng
func main() {
	// Tạo các đối tượng súng trực tiếp
	ak47 := &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}

	musket := &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}

	// In thông tin chi tiết của từng súng
	printDetails(ak47)
	printDetails(musket)
}

// Hàm in chi tiết súng
func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

// Nhược điểm của cách tiếp cận này:
// 1. Phụ Thuộc Chặt Chẽ Vào Concrete Classes:
//    - Mã client (main) cần biết rõ về các class cụ thể như Ak47 hoặc Musket.
//    - Mã dễ bị phá vỡ nếu tên class hoặc constructor thay đổi.
//
// 2. Thiếu Tính Mở Rộng:
//    - Nếu thêm loại súng mới (ví dụ: Sniper), cần sửa mã client để thêm logic.
//    - Vi phạm Nguyên tắc Open/Closed.
//
// 3. Lặp Lại Code:
//    - Logic khởi tạo các đối tượng có thể bị lặp lại ở nhiều nơi.
//
// 4. Khó Tái Sử Dụng:
//    - Không có điểm tập trung để quản lý việc khởi tạo các loại súng.
//
// 5. Khó Quản Lý Khi Có Nhiều Loại:
//    - Khi số lượng loại sản phẩm tăng lên, mã client trở nên phức tạp.
//
// 6. Không Thể Thêm Logic Khởi Tạo:
//    - Logic khởi tạo đặc biệt (như kiểm tra tính hợp lệ) phải thêm ở nhiều nơi.
