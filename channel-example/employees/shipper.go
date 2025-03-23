package employees

import (
	"fmt"
	"time"

	"channel/data"
)

type Shipper struct {
	Id int
}

func (b *Shipper) Work(transferChannel <-chan *data.MilkTea) {
	for milkTea := range transferChannel {
		fmt.Printf("Shipper %d đã nhận và đang vận chuyển %s\n", b.Id, milkTea.Name)
		time.Sleep(5 * time.Second)
		fmt.Printf("Shipper %d vận chuyển %s hoàn tất!\n", b.Id, milkTea.Name)
	}
}

type ShipperPool struct {
	shippers []*Shipper
}

func (p *ShipperPool) Init(numShipper int) {
	for idx := range numShipper {
		p.shippers = append(p.shippers, &Shipper{
			Id: idx + 1,
		})
	}
}

func (p * ShipperPool) Work(transferChannel <-chan *data.MilkTea) {
	for _, s := range p.shippers {
		go s.Work(transferChannel)
	}
}
