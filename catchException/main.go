package main

import "fmt"

func callException() error {
	return fmt.Errorf("it is error")
}

func main() {
	err := callException()
	if err != nil {
		fmt.Println(err)
	}
}
