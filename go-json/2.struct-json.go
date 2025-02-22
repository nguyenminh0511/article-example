package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Introduction struct {
	Name              string `json:"name"`
	Age               int    `json:"age"`
	Role              string `json:"role"`
	FavouriteLanguage string `json:"favourite_language"`
}

func main() {
	var myData Introduction

	// Biến fileBytes chứa dữ liệu json dưới dạng byte
	fileBytes, err := os.ReadFile("data/introduction.json")
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

	fmt.Printf("Tên của tôi là: %s\n", myData.Name)
	fmt.Printf("Hiện nay tôi %d tuổi\n", myData.Age)
	fmt.Printf("Công việc của tôi: %s\n", myData.Role)
	fmt.Printf("Ngôn ngữ lập trình yêu thích của tôi: %s\n", myData.FavouriteLanguage)

	// 5 năm sau
	myData.Age += 5
	myData.Role = "Solution Architecture"

	// Mã hóa dữ liệu
	myNewData, err := json.Marshal(&myData)
	if err != nil {
		fmt.Println("Lỗi khi mã hóa dữ liệu:", err)
		return
	}

	// Ghi lại dữ liệu vào file
	err = os.WriteFile("data/updated_introduction.json", myNewData, 0666)
	if err != nil {
		fmt.Println("Lỗi khi ghi file:", err)
		return
	}

	fmt.Println("Cập nhật thông tin thành công!")
}
