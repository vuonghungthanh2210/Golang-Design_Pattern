// evictionAlgo.go: Strategy interface
package main

import "fmt"

// Step 2: Tạo Strategy Interface
// Định nghĩa phương thức evict chung cho các thuật toán.
type EvictionAlgo interface {
	evict(c *Cache)
}

// fifo.go: Concrete strategy
// Step 3: Triển khai Concrete Strategy
// Concrete Strategy: Thuật toán FIFO.
type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

// lru.go: Concrete strategy
// Step 3: Triển khai Concrete Strategy
// Concrete Strategy: Thuật toán LRU.
type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

// lfu.go: Concrete strategy
// Step 3: Triển khai Concrete Strategy
// Concrete Strategy: Thuật toán LFU.
type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

// cache.go: Context
// Step 4: Cập nhật Context
// Lớp Cache chứa tham chiếu đến Strategy và quản lý logic chính.
type Cache struct {
	storage      map[string]string // Bộ nhớ Cache
	evictionAlgo EvictionAlgo      // Chiến lược hiện tại
	capacity     int               // Dung lượng hiện tại
	maxCapacity  int               // Dung lượng tối đa
}

// Khởi tạo Cache với chiến lược ban đầu.
// Step 4: Context khởi tạo với một chiến lược cụ thể.
func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

// Step 4: Thêm setter để thay đổi chiến lược.
func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

// Step 4: Thêm phần tử vào Cache. Nếu vượt quá dung lượng, gọi evict.
func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict() // Gọi chiến lược evict khi đầy
	}
	c.capacity++
	c.storage[key] = value
}

// Lấy phần tử từ Cache.
func (c *Cache) get(key string) {
	delete(c.storage, key)
}

// Thực hiện evict dựa trên chiến lược hiện tại.
func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// main.go: Client code
func main() {
	// Step 5: Tích hợp Client
	// Tạo chiến lược LFU ban đầu và gắn vào Cache.
	lfu := &Lfu{}
	cache := initCache(lfu)

	// Step 5: Thêm phần tử và kiểm tra chiến lược LFU.
	fmt.Println("Adding items with LFU strategy...")
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3") // Evict bằng LFU khi đầy

	// Step 5: Thay đổi chiến lược sang LRU.
	lru := &Lru{}
	cache.setEvictionAlgo(lru)
	fmt.Println("Adding items with LRU strategy...")
	cache.add("d", "4") // Evict bằng LRU khi đầy

	// Step 5: Thay đổi chiến lược sang FIFO.
	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)
	fmt.Println("Adding items with FIFO strategy...")
	cache.add("e", "5") // Evict bằng FIFO khi đầy
}

// Tóm tắt các bước:
// 1. Xác định thuật toán cần thay đổi:
//    - Các thuật toán "evict" như FIFO, LRU, LFU.
//
// 2. Tạo Strategy Interface:
//    - Interface `EvictionAlgo` với phương thức `evict`.
//
// 3. Triển khai Concrete Strategies:
//    - Các lớp `Fifo`, `Lru`, `Lfu` triển khai `EvictionAlgo`.
//
// 4. Cập nhật Context:
//    - Lớp `Cache` lưu tham chiếu đến Strategy.
//    - Thêm setter `setEvictionAlgo` để thay đổi chiến lược.
//
// 5. Tích hợp Client:
//    - Tạo các chiến lược cụ thể (`Lru`, `Fifo`, `Lfu`) và liên kết chúng với Context.
//
// Ưu điểm:
// - Thay đổi thuật toán dễ dàng trong runtime.
// - Loại bỏ logic phức tạp trong lớp chính.
// - Dễ mở rộng thêm thuật toán mới.
//
// Nhược điểm:
// - Tăng số lượng lớp và giao diện.
// - Client cần biết chi tiết các chiến lược để chọn phù hợp.
