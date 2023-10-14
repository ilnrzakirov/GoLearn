package main

import "fmt"

func main() {
point:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d %d", i, j)
			break point
		}
	}
	fmt.Println("\nEnd")
}
