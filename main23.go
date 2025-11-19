package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Открываем файл f.txt
	file, err := os.Open("f.txt")
	if err != nil {
		fmt.Println("Ошибка открытия f.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var even []int
	var odd []int

	// Считываем все числа
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err == nil {
				if n%2 == 0 {
					even = append(even, n)
				} else {
					odd = append(odd, n)
				}
			}
		}
	}

	// Сортируем: чётные по возрастанию
	sort.Ints(even)

	// Нечётные по убыванию
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	// Создаём и записываем even.txt
	evenFile, _ := os.Create("even.txt")
	defer evenFile.Close()
	for _, v := range even {
		evenFile.WriteString(fmt.Sprintf("%d ", v))
	}

	// Создаём и записываем odd.txt
	oddFile, _ := os.Create("odd.txt")
	defer oddFile.Close()
	for _, v := range odd {
		oddFile.WriteString(fmt.Sprintf("%d ", v))
	}

	fmt.Println("Готово!")
	fmt.Println("Чётные числа → even.txt")
	fmt.Println("Нечётные числа → odd.txt")
}
