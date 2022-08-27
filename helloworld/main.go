package main

import "fmt"

func main() {
	fmt.Println("Hi there, hello world!")
	add(1, 2)
}

func add(x int, y int) {
	fmt.Printf("Here is X %d Here is Y %d \n", x, y)
}
