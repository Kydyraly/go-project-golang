package main

import "fmt"

func MathOperation(a, b int, op string) float64 {
	switch op {
	case :
		return float64(a + b)
	case :
		return float64(a - b)
	case :
		return float64(a * b)
	case :
		if b != 0 {
			return float64(a) / float64(b)
		} else {
			fmt.Println("Қате: нөлге бөлуге болмайды!")
			return 0
		}
	default:
		fmt.Println("Белгісіз операция:", op)
		return 0
	}
}

func main() {
	fmt.Println("5 + 3 =", MathOperation(5, 3, "+"))
	fmt.Println("8 - 2 =", MathOperation(8, 2, "-"))
	fmt.Println("4 * 6 =", MathOperation(4, 6, "*"))
	fmt.Println("9 / 3 =", MathOperation(9, 3, "/"))
	fmt.Println("10 / 0 =", MathOperation(10, 0, "/")) 
	fmt.Println("7 ? 2 =", MathOperation(7, 2, "?"))  
}
