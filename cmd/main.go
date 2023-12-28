package main

import (
	"fmt"

	"person_api/internal/domain/person"
)

func main() {
	var p = person.Person{
		Name:     "Fabiano",
		Email:    "ugulino.fabiano@gmail.com",
		Phone:    "123",
		Document: "123456",
	}

	fmt.Println(p)
}
