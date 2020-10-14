package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"username"`           // field tag
	Password string `json:"password,omitempty"` // omitempty is to hidde the json field if it has 0 value
}

func main() {
	users := []user{
		{"Lucas", "123456"},
		{"Maiara", "123456"},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(string(out))
}
