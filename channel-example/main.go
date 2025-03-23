package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"channel/data"
	"channel/employees"
)

func main() {
	// khởi tạo
	coordinator := employees.Coordinator{}

	bartenderPool := employees.BartenderPool{}
	bartenderPool.Init(3) // khởi tạo 3 nhân viên pha chế

	shipperPool := employees.ShipperPool{}
	shipperPool.Init(2) // khởi tạo 2 shipper

	// khởi tạo các channel
	orderInputChannel := make(chan int, 5)
	defer close(orderInputChannel)
	orderTransferChannel := make(chan *data.MilkTea, 5)
	defer close(orderTransferChannel)

	// làm việc
	go bartenderPool.Work(orderInputChannel, orderTransferChannel)
	go shipperPool.Work(orderTransferChannel)
	orders := []int{1,2,3,3,3,1}
	coordinator.ReceiveOrder(orderInputChannel, orders)

	// tín hiệu kết thúc chương trình
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
	fmt.Println("exiting")
}
