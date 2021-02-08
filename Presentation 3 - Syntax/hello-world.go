package main

import (
	"fmt"
	"math"
	"runtime"
)

/****************************************
 * Functions
*****************************************/
// Functions allow you to "bundle up" certain segments of code.
// To delcare a function in golang, you just need to do the following:
// func <function name>(<parameter 1> <parameter 1 type>, <p2> <p2 type>...) <return type 1> <return type 2> {}
// The code that gets "bundled up" is inside the curly braces
func Adder(x int) int {
	// put the code here
	var y int = x + 1

	// now "return" the new integer value from the function
	return y
}

func main() {
	/****************************************
	 * "main" is the "entry point" for your program.
	 * When you run your program, the code in this area is executed line
	 * by line until it reaches the end.
	*****************************************/
	// Here, we make the program say "Hello, world!"
	fmt.Println("Hello, world!") // "print" out "Hello, world!"

	/****************************************
	 * Declaring variables
	*****************************************/
	// Here, we DECLARE a variable "x" and SET "x" to the integer value 0
	x := 0

	// You only need to use ":=" when declaring variables.
	// Once the variable is declared, you can "set" the value by using the "="
	// Add two to the integer.
	x = x + 2

	// Alternatively, the previous statement can be shortened to:
	x += 2 // x = x + 2
	x *= 2 // x = x * 2
	x -= 2 // x = x - 2
	x /= 2 // x = x / 2

	// You can also explicitly declare the variable with its data type by doing:
	// var <variable name> <data type>
	// Refer to the powerpoint for all the data types.
	// All the variables you declare in Go have to be used, otherwise the compiler
	// gives an error message. They can't just be declared.
	// Declared variables that you don't set will just go to their default values,
	// which you can look up in the documentation.
	var myString string
	myString = "Hello, world!"
	fmt.Println(myString) // "print" out the result

	// You can also declare and set on the same line as well
	var count uint64 = 429784957253

	// Here, we can print out the value of "count"
	// Notice how instead of using fmt.Println, we use fmt.Printf
	// - fmt.Println = print line
	// - fmt.Printf = print formatted
	// We use the formatted version for more complex printing.
	// Think of %v as a placeholder and write everything after the initial string.
	// Also, it is important to note that fmt.Printf does not automatically place a
	// newline character, so you need to do that yourself with the "\n" character,
	// which is called a "line feed" or a "newline" character
	fmt.Printf("%v\n", count)

	/****************************************
	 * Type Conversions
	*****************************************/
	var sqrtThis int = 4
	var f float64 = math.Sqrt(float64(sqrtThis))
	var resultInUint = uint(f)
	fmt.Println(resultInUint)

	/****************************************
	 * Control Flow
	*****************************************/
	// "if" statments allow you to enter a certain section of code
	// IF a certain condition is met. If the value is
	// is true, then it enters the section. If the value is false,
	// then it doesn't enter the section.
	// == - is equal to
	// >= - is greater than or equal to
	// <= - is less than or equal to
	// != - is not equal to
	// >  - greater than
	// <  - less than
	age := 18
	country := "USA"

	// If age is 16, then print "You are 16 years old."
	if age == 16 {
		fmt.Println("You are 16 years old.")
	} else {
		fmt.Println("You are NOT 16 years old.")
	}

	// If age is less than 21, then print "You are underage."
	if age < 21 {
		fmt.Println("You are underage.")
	} else if age == 21 {
		fmt.Println("Have a free beer!")
	}

	// AND and OR in if statements
	// If your age is greater than or equal to 21 AND your country is the USA,
	// then you are of legal age. Else, you are not.
	if age >= 21 && country == "USA" {
		fmt.Println("You are of legal age.")
	} else {
		fmt.Println("You are NOT of legal age.")
	}

	//

	// If age is NOT equal to 1, then print "You are not 1 years old.""
	if age != 1 {
		fmt.Println("You are NOT 1 years old.")
	}

	// Switch statements - Check which OS you're running on :D
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You're using a OS X!")
	case "linux":
		fmt.Println("You're using Linux!")
	default:
		fmt.Printf("%s\n", os)
	}

	// DO NOT do this lol
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")

	// DO THIS instead :D
	// FOR LOOPS
	for x := 0; x < 5; x++ {
		fmt.Println("Hello, world!")
	}

	// Golang does not have the traditional "while loops"
	// but does have unconditional for loops
	// You can instead do the following:
	running := true
	for {
		if running == false {
			fmt.Println("HELLO!")
			break // this "breaks" out of the loop
		}
		//------------------------------------------------------------------------------
		// If I uncomment the next line of code, the code will continue to go
		// on indefinitely because the keyword, "continue" just causes the loop to
		// keep going.
		//------------------------------------------------------------------------------
		// continue // THIS IS THE KEYWORD
		running = false
	}
}
