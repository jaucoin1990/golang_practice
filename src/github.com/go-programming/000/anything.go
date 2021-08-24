package main

import "fmt"

func main() {
	fmt.Println("Hello, everyone.")
	foo()
	fmt.Println("This will print out third even though it looks like it should print second.")
}

func foo() {
	fmt.Println("This is the second line.")
}