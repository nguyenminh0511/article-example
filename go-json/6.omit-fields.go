package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type UserAge int

func (a UserAge) IsZero() bool {
	return a < 0 || a >= 150
}

type UserRelationship struct {
	Name   string `json:"Name"`
	UserId int    `json:"user_id"`
}

type UserInfo struct {
	Id        int                `json:"id"`
	Name      string             `json:"name,omitempty"`
	Age       UserAge            `json:"age,omitempty,omitzero"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Lover     UserRelationship   `json:"lover,omitzero"`
	Friends   []UserRelationship `json:"friends,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitzero"`
}

func main() {
	userData := UserInfo{
		Id:       1,
		Name:     "Tuấn",
		Age:      24,
		Username: "tuan_nguyen@gmail.com",
		Password: "tuanpassword",
		Lover: UserRelationship{
			Name:   "Hoa",
			UserId: 5,
		},
		Friends: []UserRelationship{
			{
				Name:   "Hoàng",
				UserId: 6,
			},
			{
				Name: "Trung",
				UserId: 7,
			},
		},
	}

	byteData, err := json.MarshalIndent(userData, "", "    ")
	if err != nil {
		fmt.Println("Error marshal data", err)
		return
	}

	err = os.WriteFile("data/craw_data.json", byteData, 0666)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("End!")
}
