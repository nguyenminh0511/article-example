package main

import "fmt"

func main() {
	someData := []byte(`{"name": "test json"}`)
	fmt.Println("Dữ liệu dạng byte:", someData)
	fmt.Println("Dữ liệu dạng chuỗi ký tự:", string(someData))
}