package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
