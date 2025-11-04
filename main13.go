package main

import (
	"fmt"
	"time"
)

type Product struct {
	Name         string
	Price        float64
	MadeDate     time.Time
	ExpireDays   int
	Quantity     int
	Manufacturer string
}

func main() {
	products := []Product{
		{"Молоко", 450, parseDate("2025-10-20"), 10, 20, "Агрофирма А"},
		{"Хлеб", 200, parseDate("2025-10-26"), 4, 50, "Булочная №1"},
		{"Сыр", 1200, parseDate("2025-10-15"), 15, 10, "Молзавод К"},
		{"Йогурт", 300, parseDate("2025-10-24"), 6, 30, "Агрофирма А"},
		{"Масло", 950, parseDate("2025-10-10"), 20, 15, "Ферма Z"},
	}

	now := time.Now()

	fmt.Println("Товары, срок годности которых истекает через 2 дня:\n")

	found := false

	for _, p := range products {
		expireDate := p.MadeDate.AddDate(0, 0, p.ExpireDays)

		daysLeft := int(expireDate.Sub(now).Hours() / 24)

		if daysLeft == 2 {
			found = true
			fmt.Printf("Наименование: %s\n", p.Name)
			fmt.Printf("Цена: %.2f\n", p.Price)
			fmt.Printf("Дата производства: %s\n", p.MadeDate.Format("2006-01-02"))
			fmt.Printf("Срок годности (дней): %d\n", p.ExpireDays)
			fmt.Printf("Количество: %d\n", p.Quantity)
			fmt.Printf("Производитель: %s\n", p.Manufacturer)
			fmt.Printf("Дата истечения: %s\n\n", expireDate.Format("2006-01-02"))
		}
	}

	if !found {
		fmt.Println("Нет товаров, срок годности которых истекает через два дня.")
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
