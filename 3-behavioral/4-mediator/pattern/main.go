package main

import "fmt"

// Bước 1: Định nghĩa interface Train (giao diện thành phần)
// train.go: Component
type Train interface {
	arrive()        // Logic cho tàu đến
	depart()        // Logic cho tàu rời đi
	permitArrival() // Cấp phép tàu đến
}

// Bước 2: Tạo PassengerTrain làm thành phần cụ thể
// passengerTrain.go: Concrete component
type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Không thể đến, đang chờ")
		return
	}
	fmt.Println("PassengerTrain: Đã đến")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Rời đi")
	g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Được phép đến, đang đến")
	g.arrive()
}

// Bước 3: Tạo FreightTrain như một thành phần cụ thể khác
// freightTrain.go: Concrete component
type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Không thể đến, đang chờ")
		return
	}
	fmt.Println("FreightTrain: Đã đến")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Rời đi")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Được phép đến")
	g.arrive()
}

// Bước 4: Định nghĩa interface Mediator
// mediator.go: Mediator interface
type Mediator interface {
	canArrive(Train) bool  // Kiểm tra nếu tàu có thể đến
	notifyAboutDeparture() // Thông báo khi tàu rời đi
}

// Bước 5: Triển khai StationManager làm mediator cụ thể
// stationManager.go: Concrete mediator
type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true, // Ban đầu nền tảng trống
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false // Chặn nền tảng
		return true
	}
	s.trainQueue = append(s.trainQueue, t) // Đưa tàu vào hàng đợi
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true // Mở lại nền tảng
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0] // Lấy tàu đầu tiên trong hàng đợi
		s.trainQueue = s.trainQueue[1:]      // Xóa tàu khỏi hàng đợi
		firstTrainInQueue.permitArrival()    // Cho phép tàu đến
	}
}

// Bước 6: Mã Client
// main.go: Client code
func main() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive() // Tàu hành khách thử đến
	freightTrain.arrive()   // Tàu hàng thử đến
	passengerTrain.depart() // Tàu hành khách rời đi
}

// --- Tóm tắt các bước triển khai ---
// 1. Định nghĩa giao diện thành phần (Component interface): Tạo interface `Train` với các phương thức `arrive`, `depart`, và `permitArrival`.
// 2. Triển khai các thành phần cụ thể (Concrete components): Tạo `PassengerTrain` và `FreightTrain` triển khai giao diện `Train`.
// 3. Định nghĩa giao diện Mediator: Tạo interface `Mediator` để quản lý giao tiếp giữa các thành phần.
// 4. Triển khai Mediator cụ thể: Tạo lớp `StationManager` để xử lý logic cho phép tàu vào nền tảng và quản lý hàng đợi tàu.
// 5. Tích hợp trong mã Client: Sử dụng `StationManager` để điều phối giao tiếp giữa `PassengerTrain` và `FreightTrain`.

// --- Lợi ích ---
// 1. Giảm phụ thuộc: Các thành phần như tàu không giao tiếp trực tiếp, giúp mã dễ mở rộng và bảo trì.
// 2. Kiểm soát tập trung: Tất cả logic giao tiếp được xử lý bởi `StationManager`.
// 3. Tái sử dụng: Các thành phần như `PassengerTrain` hoặc `FreightTrain` có thể tái sử dụng trong bối cảnh khác.
// 4. Dễ mở rộng: Có thể thêm các loại tàu mới hoặc thay đổi logic điều phối mà không ảnh hưởng đến mã hiện tại.
