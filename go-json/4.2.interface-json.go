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
	// khai báo biến kiểu interface{}
	var unknownData interface{}

	// đọc dữ liệu từ file
	fileBytes, err := os.ReadFile("data/updated_introduction.json")
	if err != nil {
		fmt.Println("Lỗi đọc file:", err)
		return
	}

	// giải mã dữ liệu và đẩy vào biến map
	err = json.Unmarshal(fileBytes, &unknownData)
	if err != nil {
		fmt.Println("Lỗi giải mã dữ liệu:", err)
		return
	}

	switch unknownData.(type) {
    case []interface{}:
		fmt.Println("Dữ liệu có cấu trúc dạng mảng")
        // Xử lý dữ liệu với kiểu mảng 
    case map[string]interface{}:
		fmt.Println("Dữ liệu có cấu trúc dạng map")
        // Xử lý dữ liệu kiểu map
    default:
        panic("Không nhận dạng được dữ liệu!")
	}


}