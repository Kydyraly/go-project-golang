package main

import (
	"fmt"
)

type Computer struct {
	Name      string
	Frequency float64
	RAM       int
	DVD       bool
	Price     float64
}

func main() {
	computers := []Computer{
		{"Dell Inspiron", 2.5, 8, true, 450.50},
		{"HP Pavilion", 3.0, 16, false, 720.00},
		{"Lenovo ThinkPad", 2.8, 8, true, 630.25},
	}
    maxPrice := computers[0].Price
    maxComp := computers[0]
	//fmt.Println("Список компьютеров и их стоимости:")
	//var sum float64
	//for _, c := range computers {
		//fmt.Printf("%s — %.2f USD\n", c.Name, c.Price)
		//sum += c.Price
		for _, c := range computers {
		if c.Price > maxPrice {
			maxPrice = c.Price
			maxComp = c
	}

	//avg := sum / float64(len(computers))
	fmt.Printf("bolshoi summa:  %s — %.2f USD\n", maxComp.Name, maxComp.Price)
}
