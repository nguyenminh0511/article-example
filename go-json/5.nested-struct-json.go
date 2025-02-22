package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type IntroductionVer2 struct {
	Name               string    `json:"name"`
	Age                int       `json:"age"`
	Role               string    `json:"role"`
	FavouriteLanguages []string  `json:"favourite_languages"`
	CurrentCompany     Company   `json:"company"`
	Projects           []Project `json:"projects"`
}

type Company struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Project struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

func main() {
	var myData IntroductionVer2

	// Biến fileBytes chứa dữ liệu json dưới dạng byte
	fileBytes, err := os.ReadFile("data/introductionv2.json")
	if err != nil {
		fmt.Println("Lỗi khi đọc file:", err)
		return
	}

	// đưa dữ liệu vào biến myData đã khai báo trước đó
	err = json.Unmarshal(fileBytes, &myData)
	if err != nil {
		fmt.Println("Lỗi giải mã dữ liệu:", err)
		return
	}

	// fmt.Printf("Tên của tôi là: %s\n", myData.Name)
	// fmt.Printf("Hiện nay tôi %d tuổi\n", myData.Age)
	// fmt.Printf("Công việc của tôi: %s\n", myData.Role)
	// fmt.Printf("Các ngôn ngữ lập trình yêu thích của tôi: %s\n", myData.FavouriteLanguages)

	fmt.Println(myData)
}
