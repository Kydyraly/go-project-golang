package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	processNumbers()
	processText()
}

func processNumbers() {
	file, err := os.Open("number1.txt")
	if err != nil {
		fmt.Println("Ошибка открытия number1.txt:", err)
		return
	}
	defer file.Close()

	out3, _ := os.Create("number3.txt")
	defer out3.Close()

	out2, _ := os.OpenFile("number2.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer out2.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	sumMax := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			lineNum++
			continue
		}

		
		numbers := []int{}
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err == nil {
				numbers = append(numbers, n)
			}
		}

		if len(numbers) == 0 {
			lineNum++
			continue
		}

		
		maxVal := numbers[0]
		for _, v := range numbers {
			if v > maxVal {
				maxVal = v
			}
		}

		sumMax += maxVal

		
		fmt.Println("Строка", lineNum, "числа меньше первого максимального:")
		for _, v := range numbers {
			if v < maxVal {
				fmt.Println(v)
				out3.WriteString(fmt.Sprintf("%d ", v))
			}
		}
		out3.WriteString("\n")

		
		res := fmt.Sprintf("Result = %d %d\n", lineNum, maxVal)
		fmt.Print(res)
		out2.WriteString(res)

		lineNum++
	}

	fmt.Println("Сумма максимальных элементов:", sumMax)
}

func processText() {
	file, err := os.Open("text1.txt")
	if err != nil {
		fmt.Println("Ошибка открытия text1.txt:", err)
		return
	}
	defer file.Close()

	out, _ := os.Create("text2.txt")
	defer out.Close()

	scanner := bufio.NewScanner(file)
	vowels := "aeiouAEIOU"

	fmt.Println("\nСлова, начинающиеся с гласных:")

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, w := range words {
			if len(w) > 0 && isLetter(w[0]) && strings.ContainsRune(vowels, rune(w[0])) {
				fmt.Println(w)
				out.WriteString(w + "\n")
			}
		}
	}
}

func isLetter(r byte) bool {
	return unicode.IsLetter(rune(r))
}
