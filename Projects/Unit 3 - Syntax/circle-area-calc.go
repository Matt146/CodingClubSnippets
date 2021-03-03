package main

import "fmt"

func main() {
	// A = pi * r^2
	var pi float64 = 3.141592653589
	radius := 4

	var areaOfCircle float64 = pi * (float64(radius) * float64(radius))
	fmt.Println(areaOfCircle)

	number := string(500)
	fmt.Println(number)
}
