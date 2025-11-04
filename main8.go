package main

import "fmt"

func SwapValues(a, b int) (int, int) {
	return b, a
}

func main() {
	x := 10
	y := 20

	fmt.Println("бұрын: x =", x, ", y =", y)
	x, y = SwapValues(x, y)
	fmt.Println("кейін: x =", x, ", y =", y)
}
