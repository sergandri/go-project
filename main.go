package main

import (
	"fmt"
	"strings"
)

func main() {
	var income float64
	var expenses map[string]float64

	// Ввод месячных доходов
	fmt.Print("Введите ваш месячный доход: ")
	fmt.Scanln(&income)

	// Ввод расходов
	expenses = make(map[string]float64)
	for {
		var category string
		fmt.Print("Введите категорию расходов (или 'stop', чтобы закончить): ")
		fmt.Scanln(&category)
		if strings.ToLower(category) == "stop" {
			break
		}
		var amount float64
		fmt.Print("Введите сумму расходов: ")
		fmt.Scanln(&amount)
		expenses[category] = amount
	}

	// Рассчитать и вывести чистый доход
	totalExpenses := calculateTotalExpenses(expenses)
	netIncome := income - totalExpenses
	fmt.Printf("Чистый доход: %.2f\n", netIncome)

	// Проанализировать расходы
	analyzeExpenses(income, expenses)
}

func calculateTotalExpenses(expenses map[string]float64) float64 {
	total := 0.0
	for _, amount := range expenses {
		total += amount
	}
	return total
}

func analyzeExpenses(income float64, expenses map[string]float64) {
	for category, amount := range expenses {
		percentage := (amount / income) * 100
		if percentage > 30 {
			fmt.Printf("Внимание: расходы на %s составляют %.2f%% от вашего дохода. Рекомендуется снизить траты в этой категории.\n", category, percentage)
		}
	}
}
