package main

import "fmt"

func main() {
	var yearOfBirth int
	fmt.Println("Год рождения: ")
	_, err := fmt.Scanf("%d", &yearOfBirth)
	if err != nil {
		return
	}

	switch {
	case 1946 <= yearOfBirth && yearOfBirth <= 1964:
		fmt.Println("Привет, бумер!")
	case 1965 <= yearOfBirth && yearOfBirth <= 1980:
		fmt.Println("Привет, представитель X!")
	case 1981 <= yearOfBirth && yearOfBirth <= 1996:
		fmt.Println("Привет, миллениал!")
	case 1997 <= yearOfBirth && yearOfBirth <= 2012:
		fmt.Println("Привет, зумер!")
	case yearOfBirth >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Тебя нет в списке")

	}
}
