package main

import "fmt"

// CalcBMIImperial - Takes the weight in lbs, the height in inches,
// and spits out the BMI.
func CalcBMIImperial(weight float64, height float64) float64 {
	return (703 * weight) / (height * height)
}

// ClassifyBMI - Takes BMI and prints out the classification
func ClassifyBMI(bmi float64) {
	if bmi < 18.5 {
		fmt.Println("Underweight.")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Normal.")
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Overweight.")
	} else if bmi >= 30 && bmi < 400 {
		fmt.Println("Obese.")
	} else if bmi >= 400 {
		fmt.Println("Dead.")
	}
}

func main() {
	// Height & Weight inputs
	height := 68.0
	weight := 150.0

	// Calculate the BMI here
	bmi := CalcBMIImperial(weight, height)

	// Print out the BMI we calculated the 2 decimal places and classify it
	fmt.Printf("BMI: %.2f\n", bmi)
	ClassifyBMI(bmi)
}
