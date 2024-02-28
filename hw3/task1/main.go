package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите строку:")
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	characters := make([]rune, 0)

	for _, char := range input {
		characters = append(characters, char)
	}

	fmt.Println("Слайс символов строки:", characters)
}
