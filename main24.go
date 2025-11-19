package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	m   = 3 // количество строк в одной матрице
	n   = 4 // количество столбцов в одной матрице
	eps = 1e-9
)

type Matrix [m][n]float64

// Чтение всех матриц из текстового файла
func readTextMatrices(filename string) ([]Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrices []Matrix
	scanner := bufio.NewScanner(file)

	currentRow := 0
	var mat Matrix

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Пропускаем полностью пустые строки (разделители между матрицами)
		if line == "" {
			if currentRow > 0 { // если уже начали читать матрицу
				if currentRow != m {
					return nil, fmt.Errorf("матрица неполная в файле %s", filename)
				}
				matrices = append(matrices, mat)
				currentRow = 0
			}
			continue
		}

		// Читаем строку матрицы
		if currentRow >= m {
			return nil, fmt.Errorf("слишком много строк в матрице в файле %s", filename)
		}

		fields := strings.Fields(line)
		if len(fields) != n {
			return nil, fmt.Errorf("неверное количество столбцов в строке: %s", line)
		}

		for j := 0; j < n; j++ {
			val, err := strconv.ParseFloat(fields[j], 64)
			if err != nil {
				return nil, fmt.Errorf("не число: %s", fields[j])
			}
			mat[currentRow][j] = val
		}

		currentRow++
		if currentRow == m {
			matrices = append(matrices, mat)
			currentRow = 0 // готовимся к следующей матрице
		}
	}

	// Если в конце осталась незавершённая матрица
	if currentRow > 0 && currentRow != m {
		return nil, fmt.Errorf("последняя матрица неполная в файле %s", filename)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrices, nil
}

// Дописывание матриц в текстовый файл
func appendTextMatrices(filename string, matrices []Matrix) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, mat := range matrices {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if j > 0 {
					fmt.Fprint(writer, " ")
				}
				fmt.Fprintf(writer, "%.10g", mat[i][j])
			}
			fmt.Fprintln(writer)
		}
		fmt.Fprintln(writer) // пустая строка-разделитель
	}
	return writer.Flush()
}

// Сравнение матриц с погрешностью
func matricesEqual(a, b Matrix) bool {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if math.Abs(a[i][j]-b[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func containsMatrix(haystack []Matrix, needle Matrix) bool {
	for _, m := range haystack {
		if matricesEqual(m, needle) {
			return true
		}
	}
	return false
}

// Красивый вывод
func printMatrix(mat Matrix) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%12.6f ", mat[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMatrices(title string, mats []Matrix) {
	fmt.Printf("=== %s (%d матриц) ===\n", title, len(mats))
	if len(mats) == 0 {
		fmt.Println("(пусто)\n")
		return
	}
	for i, mat := range mats {
		fmt.Printf("Матрица %d:\n", i+1)
		printMatrix(mat)
	}
	fmt.Println()
}

func main() {
	file1 := "file1.txt"
	file2 := "file2.txt"

	mats1, err := readTextMatrices(file1)
	if err != nil {
		fmt.Printf("Ошибка чтения %s: %v\n", file1, err)
		return
	}
	mats2, err := readTextMatrices(file2)
	if err != nil {
		fmt.Printf("Ошибка чтения %s: %v\n", file2, err)
		return
	}

	fmt.Println("До операции:")
	printMatrices("Первый файл", mats1)
	printMatrices("Второй файл", mats2)

	var toAdd []Matrix
	for _, mat := range mats1 {
		if !containsMatrix(mats2, mat) {
			toAdd = append(toAdd, mat)
		}
	}

	if len(toAdd) == 0 {
		fmt.Println("Все матрицы из первого файла уже есть во втором — ничего не добавлено.")
		return
	}

	if err := appendTextMatrices(file2, toAdd); err != nil {
		fmt.Printf("Ошибка записи в %s: %v\n", file2, err)
		return
	}

	fmt.Printf("Добавлено %d новых матриц во второй файл.\n\n", len(toAdd))

	// Перечитываем второй файл для актуального вывода
	mats2, _ = readTextMatrices(file2)

	fmt.Println("После операции:")
	printMatrices("Первый файл", mats1)
	printMatrices("Второй файл", mats2)
}
