package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	p := &numbers

	(*p)[0] = 100

	fmt.Println("Өзгертілген тілім:", numbers)
}
