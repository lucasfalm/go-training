package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"username"`
	Password string `json:"password"`
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
