package employees

import (
	"fmt"
	"time"
)

type Coordinator struct{}

func (c *Coordinator) ReceiveOrder(orderInputChannel chan<- int, milkTeaIds []int) {
	for _, id := range milkTeaIds {
		fmt.Printf("Nhận được đơn hàng với loại đồ uống %d\n", id)
		time.Sleep(500 * time.Microsecond) // mô phỏng quá trình tạo và truyền đơn
		orderInputChannel <- id
	}
}
