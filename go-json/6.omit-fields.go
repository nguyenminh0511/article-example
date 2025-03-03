package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type UserAge int

func (a UserAge) IsZero() bool {
	return a > 150
}

type UserRelationship struct {
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}

type UserInfo struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Age       UserAge          `json:"age,omitzero"`
	Username  string           `json:"username"`
	Password  string           `json:"-"`
	Height    int              `json:"height,omitempty"`
	Weight    int              `json:"weight,omitempty"`
	Lover     UserRelationship `json:"lover,omitzero"`
	CreatedAt time.Time        `json:"created_at,omitzero"`
}

func main() {
	userData1 := UserInfo{
		Id:       1,
		Name:     "Tuấn",
		Age:      25,
		Username: "tuan_nguyen@gmail.com",
		Password: "tuanpassword",
		Height:   170,
		Weight:   71,
		Lover: UserRelationship{
			Name:   "Hoa",
			UserId: 5,
		},
		CreatedAt: time.Now(),
	}

	userData2 := UserInfo{
		Id:       2,
		Name:     "Minh",
		Age:      25,
		Username: "minh_nguyen@gmail.com",
		Password: "minhpassword",
	}

	users := []UserInfo{userData1, userData2}

	byteData, err := json.MarshalIndent(users, "", "    ")
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
