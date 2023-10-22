package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email" example:"mymail@somemail.com"`
	age   int
}

func GetPersonAge(person Person) int {
	return person.age
}

func main() {
	man := Person{
		Name:  "Alex",
		age:   22,
		Email: "mymail@somemail.com",
	}
	fmt.Println(man.Name)
	jsn, err := json.Marshal(man)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsn))
}
