package main

import "fmt"

// Bước 1: Đảm bảo rằng mô hình cốt lõi của ứng dụng có thể được biểu diễn dưới dạng cấu trúc cây.
// Ở đây, chúng ta có hai loại phần tử:
// - Leaf: Đại diện cho các tập tin đơn giản.
// - Container (Composite): Đại diện cho các thư mục, có thể chứa cả tập tin và các thư mục khác.

// Bước 2: Khai báo component interface với danh sách các phương thức có ý nghĩa cho cả các phần tử đơn giản và phức tạp.
// Component interface khai báo phương thức `search` dùng để tìm kiếm từ khóa trong cả tập tin và thư mục.
type Component interface {
	search(keyword string) // Phương thức tìm kiếm từ khóa.
}

// Bước 3: Tạo một leaf class để đại diện cho các phần tử đơn giản.
// `File` là một leaf class đại diện cho các tập tin trong hệ thống.
type File struct {
	name string // Tên của tập tin.
}

// `search` của File chỉ in thông báo tìm kiếm từ khóa trong tập tin.
func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in file '%s'\n", keyword, f.name)
}

// Bước 4: Tạo một container class để đại diện cho các phần tử phức tạp.
// `Folder` là một container class đại diện cho thư mục, có thể chứa cả tập tin và thư mục con.
type Folder struct {
	components []Component // Mảng lưu trữ các phần tử con (tập tin hoặc thư mục).
	name       string      // Tên của thư mục.
}

// `search` của Folder duyệt qua tất cả các phần tử con và ủy thác việc tìm kiếm cho chúng.
func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword '%s' in folder '%s'\n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword) // Ủy thác tìm kiếm cho các phần tử con.
	}
}

// Bước 5: Định nghĩa các phương thức để thêm và xóa các phần tử con trong container.
// Phương thức `add` thêm một phần tử (tập tin hoặc thư mục) vào thư mục.
func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func main() {
	// Tạo các leaf components (tập tin).
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	// Tạo một folder (container) và thêm file1 vào.
	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1)

	// Tạo một folder khác và thêm file2, file3 và folder1 vào.
	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	// Bắt đầu tìm kiếm từ khóa từ folder2.
	folder2.search("rose")
}

// Lợi ích của Composite Pattern
// 1. Xử lý cấu trúc cây một cách nhất quán:
// Dù là phần tử đơn giản (tập tin) hay phần tử phức tạp (thư mục), client chỉ cần làm việc thông qua một interface chung (Component).

// 2. Hỗ trợ đệ quy:
// Dễ dàng thực hiện các hành vi lặp lại trên toàn bộ cây (như tìm kiếm, tính toán, tổng hợp).

// 3. Mở rộng dễ dàng:
// Thêm phần tử mới (tập tin hoặc thư mục con) mà không ảnh hưởng đến mã hiện có.

// Khi nào nên sử dụng?
// - Khi bạn cần làm việc với một cấu trúc phân cấp (như cây thư mục).
// - Khi bạn muốn xử lý các phần tử đơn giản và phức tạp theo cách thống nhất.
// - Khi các hành vi (như tìm kiếm) cần được thực thi trên tất cả các phần tử trong cây.
