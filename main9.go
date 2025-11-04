package main

import (
	"fmt"
	"math"
)

func CircleMetrics(radius float64) (area, perimeter float64) {
	area = math.Pi * radius * radius
	perimeter = 2 * math.Pi * radius
	return area, perimeter
}

func main() {
	var r float64
	fmt.Print("радиус: ")
	fmt.Scan(&r)

	area, perimeter := CircleMetrics(r)

	fmt.Printf("Ауданы: %.2f\n", area)
	fmt.Printf("Ұзындығы: %.2f\n", perimeter)
}
