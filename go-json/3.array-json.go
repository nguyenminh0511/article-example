package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// khai báo biến dạng mảng
	var programmingLanguages []string

	// đọc dữ liệu từ file
	fileBytes, err := os.ReadFile("data/programming_languages.json")
	if err != nil {
		fmt.Println("Lỗi đọc file:", err)
		return
	}

	// giải mã dữ liệu và đẩy vào trong mảng đã khai báo trước đó
	err = json.Unmarshal(fileBytes, &programmingLanguages)
	if err != nil {
		fmt.Println("Lỗi giải mã dữ liệu:", err)
		return
	}

	fmt.Println("Một số ngôn ngữ lập trình phổ biến:")
	for i, language := range programmingLanguages {
		fmt.Printf("%d. %s\n", i, language)
	}
}