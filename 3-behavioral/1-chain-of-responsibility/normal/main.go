// Không sử dụng design pattern
package main

import "fmt"

// Patient struct lưu trữ trạng thái của bệnh nhân
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	// Khởi tạo bệnh nhân
	patient := &Patient{name: "abc"}

	// Xử lý tại lễ tân
	if !patient.registrationDone {
		fmt.Println("Lễ tân đăng ký cho bệnh nhân")
		patient.registrationDone = true
	} else {
		fmt.Println("Bệnh nhân đã được đăng ký trước đó")
	}

	// Xử lý tại bác sĩ
	if !patient.doctorCheckUpDone {
		fmt.Println("Bác sĩ đang khám cho bệnh nhân")
		patient.doctorCheckUpDone = true
	} else {
		fmt.Println("Bệnh nhân đã được bác sĩ khám trước đó")
	}

	// Xử lý tại nhà thuốc
	if !patient.medicineDone {
		fmt.Println("Nhà thuốc đang cấp thuốc cho bệnh nhân")
		patient.medicineDone = true
	} else {
		fmt.Println("Bệnh nhân đã nhận thuốc trước đó")
	}

	// Xử lý tại thu ngân
	if !patient.paymentDone {
		fmt.Println("Thu ngân đang thu tiền từ bệnh nhân")
		patient.paymentDone = true
	} else {
		fmt.Println("Bệnh nhân đã thanh toán trước đó")
	}
}

// Nhược điểm và khó khăn:
// 1. **Không thể tái sử dụng mã lệnh:**
//    - Logic xử lý của từng phòng ban (lễ tân, bác sĩ, nhà thuốc, thu ngân) không thể tái sử dụng trong các quy trình khác.
//    - Nếu có một quy trình khác cần xử lý tương tự, phải sao chép lại toàn bộ mã lệnh.

// 2. **Không linh hoạt:**
//    - Không thể thay đổi thứ tự xử lý hoặc thêm/xóa phòng ban một cách dễ dàng.
//    - Nếu cần thêm một phòng ban mới (như phòng xét nghiệm), phải sửa đổi trực tiếp hàm `main` và kiểm tra logic của từng phòng ban.

// 3. **Khó bảo trì:**
//    - Toàn bộ logic xử lý được đặt trong hàm `main`, dẫn đến mã lệnh dài và khó đọc.
//    - Khi cần sửa đổi một phần nhỏ, có nguy cơ ảnh hưởng đến toàn bộ quy trình.

// 4. **Không tuân theo nguyên tắc Single Responsibility:**
//    - Hàm `main` chịu trách nhiệm quá nhiều: từ quản lý trạng thái bệnh nhân đến xử lý từng phòng ban.

// 5. **Khó kiểm thử:**
//    - Không thể kiểm thử riêng biệt từng phòng ban vì logic xử lý được gộp chung trong một hàm lớn.
//    - Việc kiểm thử phải thực hiện trên toàn bộ quy trình, làm tăng độ phức tạp khi phát hiện lỗi.
