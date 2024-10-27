package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// triangleArea вычисляет площадь треугольника по основанию и высоте
func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

// sortArray сортирует массив целых чисел по возрастанию
func sortArray(arr []int) []int {
	sort.Ints(arr)
	return arr
}

// sumOfSquares вычисляет сумму квадратов чётных чисел от 1 до n
func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}

// isPalindrome проверяет, является ли строка палиндромом
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// isPrime проверяет, является ли число простым
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes генерирует массив простых чисел до заданного предела
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// toBinary переводит число в двоичную систему счисления
func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}
	return binary
}

// findMax находит максимальное значение в массиве
func findMax(arr []int) int {
	if len(arr) == 0 {
		panic("Массив не должен быть пустым")
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// gcd вычисляет наибольший общий делитель двух чисел
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// sumArray вычисляет сумму элементов массива
func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// getInput читает строку ввода от пользователя
func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// parseIntArray парсит строку в массив целых чисел
func parseIntArray(input string) []int {
	strArr := strings.Split(input, " ")
	var intArr []int
	for _, s := range strArr {
		num, err := strconv.Atoi(s)
		if err == nil {
			intArr = append(intArr, num)
		}
	}
	return intArr
}

func main() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Вычисление площади треугольника")
		fmt.Println("2. Сортировка массива по возрастанию")
		fmt.Println("3. Сумма квадратов чётных чисел")
		fmt.Println("4. Проверка палиндрома")
		fmt.Println("5. Проверка числа на простоту")
		fmt.Println("6. Генерация последовательности простых чисел")
		fmt.Println("7. Перевод числа в двоичную систему")
		fmt.Println("8. Нахождение максимального элемента в массиве")
		fmt.Println("9. Вычисление НОД")
		fmt.Println("10. Сумма элементов массива")
		fmt.Println("0. Выход")

		choice := getInput("Введите номер задачи: ")

		switch choice {
		case "1":
			base, _ := strconv.ParseFloat(getInput("Введите основание: "), 64)
			height, _ := strconv.ParseFloat(getInput("Введите высоту: "), 64)
			fmt.Println("Площадь треугольника:", triangleArea(base, height))
		case "2":
			arr := parseIntArray(getInput("Введите массив целых чисел через пробел: "))
			fmt.Println("Отсортированный массив:", sortArray(arr))
		case "3":
			n, _ := strconv.Atoi(getInput("Введите число n: "))
			fmt.Println("Сумма квадратов чётных чисел:", sumOfSquares(n))
		case "4":
			s := getInput("Введите строку: ")
			fmt.Println("Строка является палиндромом:", isPalindrome(s))
		case "5":
			n, _ := strconv.Atoi(getInput("Введите число: "))
			fmt.Println("Число является простым:", isPrime(n))
		case "6":
			limit, _ := strconv.Atoi(getInput("Введите предел: "))
			fmt.Println("Простые числа до", limit, ":", generatePrimes(limit))
		case "7":
			n, _ := strconv.Atoi(getInput("Введите число: "))
			fmt.Println("Число в двоичной системе:", toBinary(n))
		case "8":
			arr := parseIntArray(getInput("Введите массив целых чисел через пробел: "))
			fmt.Println("Максимальный элемент в массиве:", findMax(arr))
		case "9":
			a, _ := strconv.Atoi(getInput("Введите первое число: "))
			b, _ := strconv.Atoi(getInput("Введите второе число: "))
			fmt.Println("НОД чисел:", gcd(a, b))
		case "10":
			arr := parseIntArray(getInput("Введите массив целых чисел через пробел: "))
			fmt.Println("Сумма элементов массива:", sumArray(arr))
		case "0":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
