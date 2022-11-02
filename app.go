package main

import "fmt"

func main() {
	defer fmt.Println("Word")

	fmt.Println("hello")
}
