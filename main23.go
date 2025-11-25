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
	
	file, err := os.Open("f.txt")
	if err != nil {
		fmt.Println("Ошибка открытия f.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var even []int
	var odd []int

	
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

	
	sort.Ints(even)

	
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	
	evenFile, _ := os.Create("even.txt")
	defer evenFile.Close()
	for _, v := range even {
		evenFile.WriteString(fmt.Sprintf("%d ", v))
	}

	
	oddFile, _ := os.Create("odd.txt")
	defer oddFile.Close()
	for _, v := range odd {
		oddFile.WriteString(fmt.Sprintf("%d ", v))
	}

	fmt.Println("Готово!")
	fmt.Println("Чётные числа → even.txt")
	fmt.Println("Нечётные числа → odd.txt")
}
