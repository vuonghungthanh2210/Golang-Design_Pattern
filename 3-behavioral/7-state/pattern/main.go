// vendingMachine.go: Context
package main

import (
	"fmt"
	"log"
)

// Step 1: Xác định lớp Context
// Lớp Context quản lý trạng thái hiện tại và điều phối các hành động dựa trên trạng thái.
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

// Tạo mới một máy bán hàng với số lượng và giá mặt hàng.
func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	// Khởi tạo các trạng thái cụ thể
	hasItemState := &HasItemState{vendingMachine: v}
	itemRequestedState := &ItemRequestedState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}
	noItemState := &NoItemState{vendingMachine: v}

	// Gán trạng thái ban đầu
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v
}

// Phương thức xử lý yêu cầu các hành động
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Thay đổi trạng thái hiện tại
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Tăng số lượng mặt hàng
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount += count
}

// state.go: State Interface
// Step 2: Tạo State Interface
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// noItemState.go: Concrete State
// Step 3: Tạo Concrete State
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

// hasItemState.go: Concrete State
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

// itemRequestedState.go: Concrete State
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

// hasMoneyState.go: Concrete State
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount--
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

// main.go: Client Code
// Step 5: Tích hợp cơ chế thông báo vào Client Code
func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/*
Tóm tắt:
1. Xác định lớp Context (VendingMachine).
2. Tạo State Interface để xác định các hành vi.
3. Tạo các Concrete State để triển khai hành vi.
4. Thêm tham chiếu trạng thái vào Context.
5. Tích hợp vào Client, thay thế điều kiện bằng lời gọi trạng thái.

Lợi ích:
- Đơn giản hóa mã bằng cách chia logic theo trạng thái.
- Dễ dàng thêm hoặc chỉnh sửa trạng thái.
- Tuân thủ nguyên tắc Open/Closed Principle.
*/
