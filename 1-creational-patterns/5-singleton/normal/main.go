package main

import (
	"fmt"
)

type Single struct {
}

func NewSingle() *Single {
	fmt.Println("Creating new instance.")
	return &Single{}
}

func main() {
	// Tạo 30 instance khác nhau
	for i := 0; i < 30; i++ {
		go func() {
			instance := NewSingle()
			fmt.Printf("New Instance Address: %p\n", instance)
		}()
	}

	fmt.Scanln()
}

// Khó khăn khi không áp dụng Singleton

// 1. Không đảm bảo chỉ có một instance duy nhất:
// Trong ví dụ, 30 instance được tạo ra trong vòng lặp, dẫn đến lãng phí tài nguyên.

// 2. Truy cập không đồng nhất:
// Mỗi nơi trong chương trình sẽ nhận được một instance khác nhau, điều này gây khó khăn khi cần quản lý tài nguyên chung, như cơ sở dữ liệu, file, hoặc cấu hình ứng dụng.

// 3. Lãng phí tài nguyên:
// Mỗi lần tạo một instance mới tiêu tốn bộ nhớ và CPU, đặc biệt khi đối tượng nặng (heavy object) như kết nối cơ sở dữ liệu hoặc kết nối mạng.

// 4. Thiếu an toàn trong môi trường đa luồng:
// Không có cơ chế kiểm soát truy cập giữa các luồng. Trong ví dụ trên, nếu 30 luồng cùng tạo instance, có khả năng dẫn đến lỗi cạnh tranh (race condition) khi truy cập tài nguyên.

// 5. Khó bảo trì và mở rộng:
// Nếu cần thay đổi logic sao chép hoặc quản lý instance, bạn phải sửa đổi ở nhiều nơi thay vì một điểm tập trung như trong Singleton.
