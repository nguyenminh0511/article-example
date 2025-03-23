package data

type MilkTea struct {
	Id             int
	Name           string
	ProcessingTime int
}

func MockData() map[int]MilkTea {
	menu := make(map[int]MilkTea)
	menu[1] = MilkTea{
		Id:             1,
		Name:           "Trà sữa chân châu",
		ProcessingTime: 5,
	}
	menu[2] = MilkTea{
		Id:             2,
		Name:           "Hồng trà sữa",
		ProcessingTime: 4,
	}
	menu[3] = MilkTea{
		Id:             3,
		Name:           "Matcha kem cheese",
		ProcessingTime: 6,
	}
	return menu
}
