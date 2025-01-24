package main

import "fmt"

// IOtp defines the interface for OTP operations.
// 1. Phân tích thuật toán, xác định các bước chung và riêng.
// otp.go: Template method
type IOtp interface {
	genRandomOTP(int) string       // Bước 1: Tạo OTP ngẫu nhiên.
	saveOTPCache(string)           // Bước 2: Lưu OTP vào bộ nhớ đệm.
	getMessage(string) string      // Bước 3: Chuẩn bị nội dung thông báo.
	sendNotification(string) error // Bước 4: Gửi thông báo OTP.
}

// Otp provides the template method.
// 2. Tạo lớp cha (abstract class) với:
// Template method: Xác định khung sườn thuật toán.
type Otp struct {
	iOtp IOtp // Phụ thuộc vào lớp con thông qua interface.
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)   // Tạo OTP.
	o.iOtp.saveOTPCache(otp)                // Lưu OTP vào cache.
	message := o.iOtp.getMessage(otp)       // Chuẩn bị nội dung thông báo.
	err := o.iOtp.sendNotification(message) // Gửi thông báo.
	if err != nil {
		return err
	}
	return nil
}

// sms.go: Concrete implementation
// Sms implements the IOtp interface for SMS OTP.
type Sms struct {
	Otp
}

// 4. Tạo các lớp con cụ thể, triển khai các bước trừu tượng.
func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

// email.go: Concrete implementation
// Email implements the IOtp interface for email OTP.
type Email struct {
	Otp
}

// 4. Tạo các lớp con cụ thể, triển khai các bước trừu tượng.
func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

// main.go: Client code
func main() {
	// Tạo đối tượng SMS OTP và thực thi quy trình.
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")

	// Tạo đối tượng Email OTP và thực thi quy trình.
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}

/*
Tóm tắt các bước thực hiện:
1. Phân tích thuật toán, xác định các bước chung và riêng:
   - Các bước chung: Tạo OTP, lưu cache, chuẩn bị nội dung, gửi thông báo.
   - Các bước riêng: Cách tạo OTP và cách gửi thông báo (SMS, Email).

2. Tạo lớp cha (Otp) chứa template method `genAndSendOTP` để thực hiện khung sườn của thuật toán.

3. Tạo interface IOtp định nghĩa các phương thức cần thiết cho các bước cụ thể.

4. Tạo các lớp con (Sms, Email) triển khai interface IOtp để cung cấp cách thực hiện cho từng bước.

5. Tích hợp client sử dụng template method để thực hiện quy trình OTP tùy thuộc vào loại (SMS, Email).
*/
