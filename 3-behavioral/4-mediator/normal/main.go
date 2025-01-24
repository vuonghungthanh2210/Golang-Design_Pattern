package main

import "fmt"

// PassengerTrain đại diện cho tàu chở khách
type PassengerTrain struct {
	isPlatformFree *bool
	trainQueue     *[]*PassengerTrain
}

func (g *PassengerTrain) arrive() {
	if *g.isPlatformFree {
		*g.isPlatformFree = false
		fmt.Println("PassengerTrain: Arrived")
	} else {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		*g.trainQueue = append(*g.trainQueue, g)
	}
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	*g.isPlatformFree = true
	// Thông báo tàu tiếp theo
	if len(*g.trainQueue) > 0 {
		nextTrain := (*g.trainQueue)[0]
		*g.trainQueue = (*g.trainQueue)[1:]
		nextTrain.arrive()
	}
}

// FreightTrain đại diện cho tàu hàng
type FreightTrain struct {
	isPlatformFree *bool
	trainQueue     *[]*FreightTrain
}

func (g *FreightTrain) arrive() {
	if *g.isPlatformFree {
		*g.isPlatformFree = false
		fmt.Println("FreightTrain: Arrived")
	} else {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		*g.trainQueue = append(*g.trainQueue, g)
	}
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	*g.isPlatformFree = true
	// Thông báo tàu tiếp theo
	if len(*g.trainQueue) > 0 {
		nextTrain := (*g.trainQueue)[0]
		*g.trainQueue = (*g.trainQueue)[1:]
		nextTrain.arrive()
	}
}

func main() {
	isPlatformFree := true
	passengerQueue := make([]*PassengerTrain, 0)
	freightQueue := make([]*FreightTrain, 0)

	passengerTrain := &PassengerTrain{
		isPlatformFree: &isPlatformFree,
		trainQueue:     &passengerQueue,
	}

	freightTrain := &FreightTrain{
		isPlatformFree: &isPlatformFree,
		trainQueue:     &freightQueue,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

// Nhược Điểm và Khó Khăn khi không sử dụng Mediator Pattern
//
// 1. Phụ thuộc chồng chéo:
//    - Từng loại tàu phải tự quản lý trạng thái nền tảng và hàng đợi của riêng mình.
//    - Gây khó khăn trong việc bảo trì và làm tăng sự phụ thuộc giữa các tàu.
//
// 2. Lặp lại mã:
//    - Logic kiểm tra trạng thái nền tảng và thông báo tàu tiếp theo được lặp lại trong mỗi loại tàu.
//
// 3. Khả năng mở rộng kém:
//    - Nếu cần thêm loại tàu mới, bạn phải sao chép và sửa đổi logic, dễ gây ra lỗi.
//
// 4. Thiếu tính linh hoạt:
//    - Không thể dễ dàng thay đổi cách thức điều phối, ví dụ ưu tiên loại tàu này hơn tàu khác.
//
// 5. Khó tái sử dụng:
//    - Các lớp tàu bị ràng buộc chặt chẽ với logic quản lý nền tảng, không thể tái sử dụng trong các ngữ cảnh khác.
//
// 6. Tăng độ phức tạp:
//    - Mã trở nên khó đọc và bảo trì khi logic điều phối phức tạp hơn.
