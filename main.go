package main

import "fmt"

func presentation() {
	fmt.Println("Check Website Availability")
	fmt.Println("..........................")

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
	fmt.Println()
	fmt.Println()
}

func selectOption() int8 {
	var input int8
	fmt.Print("Select an option: ")
	fmt.Scanf("%d", &input)

	return input
}

func handleCommand(input int) {
	switch input {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Invalid input. Please select a number between 1 and 3")
	}
}

func main() {
	presentation()

	var userInput = selectOption()
	handleCommand(int(userInput))
}
