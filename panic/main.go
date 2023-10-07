package main

import "fmt"

func div(a, b int) int {
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	div(5, 0)
	fmt.Println("don't you see")
}
