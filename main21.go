package main

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Қате [%d]: %s", e.Code, e.Message)
}
func findUser(username string) error {
	if username == "" {
		return MyError{404, "Пайдаланушы табылмады"}
	}
	if username == "admin" {
		return MyError{500, "Серверлік қате пайда болды"}
	}
	return nil
}
func main() {
	users := []string{"", "admin", "user1"}

	for _, u := range users {
		err := findUser(u)
		if err != nil {
			fmt.Println("Қате:", err)
		} else {
			fmt.Println("Пайдаланушы табылды:", u)
		}
	}
}
