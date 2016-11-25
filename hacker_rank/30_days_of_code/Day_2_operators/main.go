package main

import "fmt"

type mealFinance struct {
	mealCost   float64
	tipPercent float64
	taxPercent float64
	totalCost  int
}

func main() {
	meal := mealFinance{}

	fmt.Scanf("%v\n%v\n%v", &meal.mealCost, &meal.tipPercent, &meal.taxPercent)

	tip := meal.calculateTip()
	tax := meal.calculateTax()
	totalCost := meal.calculateTotalCost(tip, tax)
	roundedTotalCost := calculateRoundedTotalCost(totalCost)

	fmt.Printf("The total meal cost is %d dollars.\n", roundedTotalCost)
}

func (m mealFinance) calculateTip() float64 {
	return m.mealCost * (m.tipPercent / 100)
}

func (m mealFinance) calculateTax() float64 {
	return m.mealCost * (m.taxPercent / 100)
}

func (m mealFinance) calculateTotalCost(tax float64, tip float64) float64 {
	return m.mealCost + tax + tip
}

func calculateRoundedTotalCost(totalCost float64) int {
	if shouldRoundUp(totalCost) {
		return roundUp(totalCost)
	}
	return roundDown(totalCost)
}

func shouldRoundUp(totalCost float64) bool {
	if (int((totalCost - float64(int(totalCost))) * 10)) < 5 {
		return false
	}
	return true
}

func roundUp(totalCost float64) int {
	return int(totalCost + 1)
}

func roundDown(totalCost float64) int {
	return int(totalCost)
}
