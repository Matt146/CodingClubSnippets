package main

import (
	"bufio" // buffered I/O
	"fmt"
	"os"
)

func main() {
	// Inputs
	// os.Stdin = standard input (input into the command line)
	// os.Stdout = standard output (output to the command line)
	// os.Stderr = standard error (output errors to the command line)
	fmt.Println("Please type your name:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
