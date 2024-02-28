package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Age    int
	Grades []int
}

func main() {
	var students []Student

	fmt.Println("Введите информацию о студентах:")

	for {
		student := readStudent()
		students = append(students, student)

		fmt.Print("Хотите добавить еще одного студента? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" {
			break
		}
	}

	printStudents(students)
}

func readStudent() Student {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя студента: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Введите возраст студента: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	age, _ := strconv.Atoi(ageInput)

	fmt.Print("Введите оценки студента (через запятую, без пробелов): ")
	gradesInput, _ := reader.ReadString('\n')
	gradesInput = strings.TrimSpace(gradesInput)
	gradesStr := strings.Split(gradesInput, ",")
	grades := make([]int, len(gradesStr))
	for i, gradeStr := range gradesStr {
		grades[i], _ = strconv.Atoi(gradeStr)
	}

	return Student{Name: name, Age: age, Grades: grades}
}

func printStudents(students []Student) {
	fmt.Println("Список студентов:")
	for _, student := range students {
		fmt.Printf("Имя: %s, Возраст: %d, Оценки: %v\n", student.Name, student.Age, student.Grades)
	}
}
