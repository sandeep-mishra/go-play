package main

import "fmt"

func main() {

	functions := map[int]string{
		1: "hello",
		2: "play_array",
	}

	var choice int
	choice = choose(functions)

	switch choice {
	case 1:
		hello()
	case 2:
		play_array()
	}

	fmt.Println("Exiting")

}

func choose(functions map[int]string) int {

	var choice int
	fmt.Println("Please choose function to execute: ")
	fmt.Println(functions)
	fmt.Scanf("%d\n", &choice)
	return choice
}

func hello() {
	var input string
	fmt.Println("Please enter your name : ")
	fmt.Scanf("%s\n", &input)
	fmt.Println("Hello", input)
}

func play_array() {
	var arr [10]int
	for _, v := range arr {
		fmt.Printf("%d", v)
	}

	var val int
	for index, _ := range arr {
		arr[index] = val + 1
	}

	fmt.Println("\nAfter updates : ")
	for _, v := range arr {
		fmt.Printf("%d", v)
	}
	fmt.Println()

}
