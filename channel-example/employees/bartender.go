package employees

import (
	"fmt"
	"time"

	"channel/data"
)

type Bartender struct {
	Id int
}

func (b *Bartender) MakeMilkTea(milkTea *data.MilkTea, transferChannel chan<- *data.MilkTea) {
	fmt.Printf("Nhân viên pha chế %d: Nhận được đơn hàng làm %s, đang chế biến...\n", b.Id, milkTea.Name)
	time.Sleep(time.Duration(milkTea.ProcessingTime) * time.Second) // Mô phỏng quá trình pha chế
	transferChannel <- milkTea
}

func (b *Bartender) Work(orderInputChannel <-chan int, transferChannel chan<- *data.MilkTea, menu map[int]data.MilkTea) {
	for milkTeaId := range orderInputChannel {
		milkTea, ok := menu[milkTeaId]
		if !ok {
			fmt.Printf("Nhân viên pha chế %d: Đồ uống với ID %d không nằm trong danh sách menu!\n", b.Id, milkTeaId)
			continue
		}
		b.MakeMilkTea(&milkTea, transferChannel)
	}
}

type BartenderPool struct {
	bartenders []*Bartender
	Menu       map[int]data.MilkTea
}

func (p *BartenderPool) initMenu() {
	p.Menu = data.MockData()
}

func (p *BartenderPool) initPool(numBartender int) {
	for idx := range numBartender {
		p.bartenders = append(p.bartenders, &Bartender{
			Id: idx + 1,
		})
	}
}

func (p *BartenderPool) Init(numBartender int) {
	p.initPool(numBartender)
	p.initMenu()
}

func (p *BartenderPool) Work(orderInputChannel <-chan int, transferChannel chan<- *data.MilkTea) {
	for _, b := range p.bartenders {
		go b.Work(orderInputChannel, transferChannel, p.Menu)
	}
}
