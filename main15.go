package main

import "fmt"

type Student struct {
	name  string
	grade int
}

func main() {
	s := Student{name: "Қыдырәлі", grade: 85}

	p := &s

	p.grade = 95
	p.name = "Саламатұлы Қыдырәлі"

	fmt.Println("Студент аты:", s.name)
	fmt.Println("Бағасы:", s.grade)
}
