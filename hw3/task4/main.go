package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	courseCounts := make(map[string]int)

	fmt.Println("Введите информацию о студентах (для завершения введите 'exit'):")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите имя студента: ")
		scanner.Scan()
		name := scanner.Text()
		if name == "exit" {
			break
		}

		fmt.Print("Введите курс студента: ")
		scanner.Scan()
		course := scanner.Text()

		courseCounts[course]++
	}

	fmt.Println("Общее количество студентов по каждому курсу:")
	for course, count := range courseCounts {
		fmt.Printf("Курс %s: %d студент(ов)\n", course, count)
	}
}
