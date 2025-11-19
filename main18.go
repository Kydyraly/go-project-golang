package main

import (
	"fmt"
	"strconv"
)
func main() {
	var input string
	fmt.Print("Мән енгізіңіз: ")
	fmt.Scan(&input)

	if _, err := strconv.Atoi(input); err == nil {
		fmt.Println("int типі")
		return
	}
	if _, err := strconv.ParseFloat(input, 64); err == nil {
		fmt.Println("float типі")
		return
	}

	fmt.Println("string типі")
}
