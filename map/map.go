package main

import "fmt"

func main() {
	products := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var oreder_sum int
	for _, v := range order {
		oreder_sum += products[v]
	}
	fmt.Println("Заказ на сумму", oreder_sum)
}
