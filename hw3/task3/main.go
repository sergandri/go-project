package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	translations := make(map[string]string)

	fmt.Println("Введите переводы слов (для завершения введите 'exit'):")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите слово на первом языке: ")
		scanner.Scan()
		word := scanner.Text()
		if word == "exit" {
			break
		}

		fmt.Print("Введите его перевод на втором языке: ")
		scanner.Scan()
		translation := scanner.Text()

		translations[word] = translation
	}

	for {
		fmt.Print("Введите слово для поиска его перевода (для выхода введите 'exit'): ")
		scanner.Scan()
		searchWord := scanner.Text()
		if searchWord == "exit" {
			break
		}

		translation, found := translations[searchWord]
		if found {
			fmt.Printf("Перевод слова '%s': %s\n", searchWord, translation)
		} else {
			fmt.Printf("Перевод для слова '%s' не найден\n", searchWord)
		}
	}
}
