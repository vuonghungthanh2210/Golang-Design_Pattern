package main

import "fmt"

// AdidasShoe: Concrete Product
type AdidasShoe struct {
	logo string
	size int
}

// NikeShoe: Concrete Product
type NikeShoe struct {
	logo string
	size int
}

// AdidasShirt: Concrete Product
type AdidasShirt struct {
	logo string
	size int
}

// NikeShirt: Concrete Product
type NikeShirt struct {
	logo string
	size int
}

func main() {
	// Khởi tạo giày và áo thun cho Adidas
	adidasShoe := AdidasShoe{
		logo: "adidas",
		size: 14,
	}

	adidasShirt := AdidasShirt{
		logo: "adidas",
		size: 14,
	}

	// Khởi tạo giày và áo thun cho Nike
	nikeShoe := NikeShoe{
		logo: "nike",
		size: 14,
	}

	nikeShirt := NikeShirt{
		logo: "nike",
		size: 14,
	}

	// In thông tin chi tiết
	printShoeDetails(adidasShoe.logo, adidasShoe.size)
	printShirtDetails(adidasShirt.logo, adidasShirt.size)

	printShoeDetails(nikeShoe.logo, nikeShoe.size)
	printShirtDetails(nikeShirt.logo, nikeShirt.size)
}

func printShoeDetails(logo string, size int) {
	fmt.Printf("Logo: %s\n", logo)
	fmt.Printf("Size: %d\n", size)
}

func printShirtDetails(logo string, size int) {
	fmt.Printf("Logo: %s\n", logo)
	fmt.Printf("Size: %d\n", size)
}

/*
Những khó khăn và nhược điểm:

1. Mã phụ thuộc chặt chẽ vào các class cụ thể:
   - Mỗi lần thêm một loại sản phẩm hoặc thương hiệu mới, bạn phải sửa đổi mã client, dẫn đến sự phụ thuộc chặt chẽ vào các class cụ thể.
   - Vi phạm nguyên tắc Open/Closed Principle.

2. Thiếu tính mở rộng:
   - Nếu có một thương hiệu hoặc sản phẩm mới, mã client sẽ phải cập nhật toàn bộ logic khởi tạo và sử dụng.

3. Mã trùng lặp:
   - Khởi tạo sản phẩm có thể lặp đi lặp lại tại nhiều nơi trong mã client, dẫn đến mã khó bảo trì.

4. Khó đảm bảo tính đồng bộ:
   - Không có cơ chế đảm bảo rằng giày và áo thun của cùng một bộ phải thuộc cùng một thương hiệu.

5. Khó kiểm tra và bảo trì:
   - Khi mã client quá phụ thuộc vào logic khởi tạo cụ thể, việc kiểm tra hoặc sửa đổi sẽ trở nên khó khăn.

6. Tăng độ phức tạp khi mở rộng:
   - Việc thêm các biến thể sản phẩm hoặc nhóm sản phẩm sẽ dẫn đến sự gia tăng độ phức tạp trong mã client.

Giải pháp tốt hơn: Sử dụng Abstract Factory Pattern để tách biệt logic khởi tạo sản phẩm khỏi mã client và đảm bảo tính đồng bộ giữa các sản phẩm trong cùng một bộ.
*/
