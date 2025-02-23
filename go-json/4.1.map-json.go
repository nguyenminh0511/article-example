package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// khai báo biến kiểu map
	var unknownKeyValueData map[string]interface{}

	// đọc dữ liệu từ file
	fileBytes, err := os.ReadFile("data/updated_introduction.json")
	if err != nil {
		fmt.Println("Lỗi đọc file:", err)
		return
	}

	// giải mã dữ liệu và đẩy vào biến map
	err = json.Unmarshal(fileBytes, &unknownKeyValueData)
	if err != nil {
		fmt.Println("Lỗi giải mã dữ liệu:", err)
		return
	}

	// In dữ liệu ra màn hình
	fmt.Println("Tôi quên mất nội dung file rồi, hãy ghi ra cho tôi nhé!")
	for k, v := range unknownKeyValueData {
		fmt.Printf("- %s: %v\n", k, v)
	}
}