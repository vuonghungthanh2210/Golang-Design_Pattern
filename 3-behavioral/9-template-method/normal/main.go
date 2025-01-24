// otp.go: Không sử dụng Template Method Pattern
package main

import (
	"errors"
	"fmt"
)

type OtpService struct{}

func (o *OtpService) sendOtp(channel string, otpLength int) error {
	var otp string
	var message string

	// Bước 1: Tạo OTP
	otp = o.generateOtp(otpLength)

	// Bước 2: Lưu OTP vào bộ nhớ đệm
	err := o.saveOtpCache(otp, channel)
	if err != nil {
		return err
	}

	// Bước 3: Chuẩn bị nội dung
	if channel == "SMS" {
		message = fmt.Sprintf("SMS OTP: %s", otp)
	} else if channel == "Email" {
		message = fmt.Sprintf("EMAIL OTP: %s", otp)
	} else {
		return errors.New("Unsupported channel")
	}

	// Bước 4: Gửi OTP
	err = o.sendNotification(channel, message)
	if err != nil {
		return err
	}

	return nil
}

func (o *OtpService) generateOtp(length int) string {
	// Tạo mã OTP ngẫu nhiên (giả sử là "1234" cho đơn giản)
	return "1234"
}

func (o *OtpService) saveOtpCache(otp string, channel string) error {
	// Lưu OTP vào bộ nhớ đệm (chỉ giả lập việc lưu)
	fmt.Printf("Saving OTP %s for channel %s\n", otp, channel)
	return nil
}

func (o *OtpService) sendNotification(channel string, message string) error {
	if channel == "SMS" {
		fmt.Printf("Sending SMS: %s\n", message)
	} else if channel == "Email" {
		fmt.Printf("Sending Email: %s\n", message)
	} else {
		return errors.New("Unsupported notification channel")
	}
	return nil
}

// main.go: Client code
func main() {
	otpService := &OtpService{}

	fmt.Println("Sending OTP via SMS:")
	err := otpService.sendOtp("SMS", 4)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println("\nSending OTP via Email:")
	err = otpService.sendOtp("Email", 4)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

// Nhược Điểm và Khó Khăn (khi không sử dụng Design Pattern)

// 1. Mã Nguồn Phức Tạp và Cồng Kềnh:
//    - Hàm `sendOtp` xử lý toàn bộ quy trình OTP cho mọi kênh (SMS, Email). Điều này làm tăng độ phức tạp, đặc biệt khi thêm nhiều kênh thông báo mới.

// 2. Khó Mở Rộng:
//    - Mỗi khi cần thêm kênh mới (ví dụ: WhatsApp hoặc Push Notification), phải sửa đổi hàm `sendOtp`, dẫn đến vi phạm nguyên tắc Open/Closed Principle (OCP).

// 3. Trùng Lặp Logic:
//    - Các bước như tạo OTP, lưu vào bộ nhớ đệm và chuẩn bị thông báo có thể giống nhau giữa các kênh, nhưng bị lặp lại nhiều lần, gây dư thừa mã nguồn.

// 4. Khó Bảo Trì:
//    - Việc thay đổi logic của một bước (ví dụ: thay đổi cách tạo OTP) yêu cầu phải kiểm tra toàn bộ mã nguồn để đảm bảo không bỏ sót phần nào, tăng nguy cơ gây lỗi.

// 5. Thiếu Tính Linh Hoạt:
//    - Không thể thay đổi một bước cụ thể trong quy trình OTP mà không ảnh hưởng đến các bước khác. Ví dụ: nếu muốn thay đổi cách gửi thông báo cho từng kênh, sẽ phải chỉnh sửa logic bên trong cùng một hàm.

// 6. Khó Tái Sử Dụng:
//    - Nếu cần sử dụng lại chỉ một phần quy trình (ví dụ: gửi thông báo mà không tạo OTP), không thể tách riêng được các bước một cách dễ dàng.

// 7. Phụ Thuộc Cao:
//    - Tất cả các kênh phụ thuộc chặt chẽ vào một hàm duy nhất, gây khó khăn khi debug hoặc nâng cấp một kênh cụ thể mà không ảnh hưởng đến các kênh khác.

// Kết Luận:
// - Sử dụng Template Method Pattern sẽ giúp tách biệt rõ ràng các bước trong quy trình OTP, cải thiện tính mở rộng, linh hoạt, và dễ bảo trì.
