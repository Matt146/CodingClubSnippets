package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
		fmt.Println("Aren")
	*/
	/*
			for x := 0; x < 10; x++ {
				fmt.Println("Aren")
			}
		/*
			running = true;
			while (running) {
				printf("Running...")
			}
	*/
	/*
		running := true
		for {
			if running != true {
				break
			}
			fmt.Println("Aren")
			continue
			running = false
		}

		// land over
		fmt.Println("Exited the loop.")
	*/
	fmt.Println("Please type your number:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// Text to integer
	if len(text) > 2 {
		text = text[0 : len(text)-2]
	} else {
		panic("Invalid string")
	}
	textNum, err := strconv.Atoi(text)
	if err != nil {
		panic("Error! Could not convert from string to int!")
	}

	if textNum == 10 {
		fmt.Println("Congratulations! You guessed the correct number!")
	}
}
