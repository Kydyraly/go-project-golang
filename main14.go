package main

import (
	"fmt"
	"time"
)

type Car struct {
	Brand        string
	Manufacturer string
	Type         string
	Year         int
	RegDate      time.Time
}

func main() {
	cars := []Car{
		{"Toyota Corolla", "Toyota", "седан", 1998, parseDate("2025-03-15")},
		{"Lada Niva", "АвтоВАЗ", "внедорожник", 1995, parseDate("2025-02-20")},
		{"BMW X5", "BMW", "кроссовер", 2005, parseDate("2024-09-01")},
		{"Ford Focus", "Ford", "седан", 1999, parseDate("2025-01-10")},
		{"Mercedes E200", "Mercedes-Benz", "седан", 1998, parseDate("2023-11-30")},
	}

	now := time.Now()
	fmt.Println("Машины, произведённые до 2000 года и зарегистрированные менее года назад:\n")

	found := false

	for _, c := range cars {
		if c.Year < 2000 {
			monthsAgo := now.Sub(c.RegDate).Hours() / 24 / 30
			if monthsAgo < 12 {
				found = true
				fmt.Printf("Марка: %s\n", c.Brand)
				fmt.Printf("Производитель: %s\n", c.Manufacturer)
				fmt.Printf("Тип: %s\n", c.Type)
				fmt.Printf("Год выпуска: %d\n", c.Year)
				fmt.Printf("Дата регистрации: %s\n", c.RegDate.Format("2006-01-02"))
				fmt.Println()
			}
		}
	}

	if !found {
		fmt.Println("Нет машин, удовлетворяющих условиям.")
	}
}

func parseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Ошибка в дате:", dateStr)
		return time.Now()
	}
	return t
}
