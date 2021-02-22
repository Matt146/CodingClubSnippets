package main

import "fmt"

func main() {
	// Main inputs
	name := "Michael Jordan"
	team := "Wizards, Bulls"
	age := 58

	// Calculate free throw percentage
	freeThrowsAttempted := 8772
	freeThrowsMade := 7327
	var freeThrowPercentage float64 = float64(freeThrowsMade) / float64(freeThrowsAttempted)

	// Calculate two pointer percentage
	twoPointersAttempted := 22759
	twoPointersMade := 11611
	var twoPointerPercentage float64 = float64(twoPointersMade) / float64(twoPointersAttempted)

	// Calculate three pointer percentage
	threePointersAttempted := 1778
	threePointersMade := 581
	var threePointersPercentage float64 = float64(threePointersMade) / float64(threePointersAttempted)

	// Print out the results
	fmt.Printf("Player name: %v\n", name)
	fmt.Printf("Team: %v\n", team)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Free throw percentage: %.2f\n", freeThrowPercentage*100)
	fmt.Printf("Two pointer percentage: %.2f\n", twoPointerPercentage*100)
	fmt.Printf("Three pointer percentage: %.2f\n", threePointersPercentage*100)
}
