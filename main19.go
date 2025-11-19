package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func detectType(input string) (string, error) {
	input = strings.TrimSpace(input) 

	if input == "" {
		return "", errors.New("мән енгізілмеген")
	}

	if _, err := strconv.Atoi(input); err == nil {
		return "int типі", nil
	}

	if _, err := strconv.ParseFloat(input, 64); err == nil {
		return "float типі", nil
	}

	return "string типі", nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Мән енгізіңіз: ")

	input, _ := reader.ReadString('\n') 

	result, err := detectType(input)
	if err != nil {
		fmt.Println("Қате:", err)
	} else {
		fmt.Println("Нәтиже:", result)
	}
}
