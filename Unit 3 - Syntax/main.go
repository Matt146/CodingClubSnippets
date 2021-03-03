package main

import "fmt"

// Write a function that prints "Hello, world" for you
// and returns the value 69.
func HelloWorldFunction() int {
	fmt.Println("Hello, world!")

	return 69
}

func HelloWorldFunction1(myFavNum int) (int, int) {
	fmt.Println(myFavNum)
	fmt.Println("Hello, world!")

	return 69, 420
}

func AddOneToThisFloat(floatValue float64) float64 {
	return floatValue + 1
}

func main() {
	ticketID := 67456878
	winningNumber := 42898942
	winningNumber2 := 67456878
	secondPrizeNumber := 2489891

	x := 0
	fmt.Println(x)
	x = 1

	if ticketID == winningNumber || ticketID == winningNumber2 {
		// execute the code here
		fmt.Println("CONGRATULATIONS! YOU WON THE LOTTERY!")
	} else if ticketID == secondPrizeNumber {
		fmt.Println("CONGRATULATIONS! You won the second prize!")
	} else {
		fmt.Println("You lost :(")
	}
}
