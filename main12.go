package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	LastName  string
	BirthYear int
	Position  string
	Salary    float64
	Education string
}

func main() {
	employees := []Employee{
		{"Халелов", 1985, "инженер", 250000, "высшее"},
		{"Мансуров", 1990, "менеджер", 200000, "высшее"},
		{"Дүйсенбайқызы", 1988, "инженер", 270000, "высшее"},
		{"Тұрақбай", 1995, "техник", 180000, "среднее специальное"},
		{"Саламатұлы", 1983, "инженер", 300000, "высшее"},
	}

	count := 0

	fmt.Println("Список работников-инженеров:\n")

	for _, e := range employees {
		if strings.ToLower(e.Position) == "инженер" {
			count++
			fmt.Printf("Фамилия: %s\n", e.LastName)
			fmt.Printf("Год рождения: %d\n", e.BirthYear)
			fmt.Printf("Должность: %s\n", e.Position)
			fmt.Printf("Зарплата: %.2f\n", e.Salary)
			fmt.Printf("Образование: %s\n\n", e.Education)
		}
	}

	fmt.Printf("Количество инженеров: %d\n", count)
}
