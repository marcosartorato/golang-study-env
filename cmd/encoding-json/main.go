package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	/*
		Marshal Go struct to JSON
	*/
	fmt.Println("Marshal example")
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{Name: "Alice", Age: 30}
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", string(data))

	/*
		Unmarshal JSON to Go struct
	*/
	fmt.Println("\nUnmarshal example")
	var p2 Person
	err = json.Unmarshal(data, &p2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Go struct: %+v\n", p2)

	/*
		Validate JSON

		Validation can be done in two ways:
		1. Basic validation using json.Valid
		2. Advanced validation using struct tags and a validation library
	*/
	fmt.Println("\nBasic validation example. See code for details.")
	tests := []string{
		`{"name":"Bob","age":25}`,
		`{"name":"Charlie","age":"}`,
	}
	for _, str := range tests {
		fmt.Printf("Basic validation passed for %s: %t\n", str, json.Valid([]byte(str)))
	}

	// Test validator tags
	fmt.Println("\nAdvanced validation example. See code for details.")
	type PersonDetail struct {
		FirstName  string `json:"first_name" validate:"required"`
		MiddleName string `json:"middle_name" validate:"omitempty"` // Not all of us have middle names.
		LastName   string `json:"last_name" validate:"required"`
		Age        int    `json:"age" validate:"required,min=0"` // Age cannot be negative.
	}
	tests = []string{
		`{"first_name":"John","last_name":"Doe","age":25}`,
		`{"first_name":"John","age":-5}`,
	}
	v := validator.New()
	for _, str := range tests {
		var pd PersonDetail
		if err = json.Unmarshal([]byte(str), &pd); err != nil {
			fmt.Printf("Unmarshal failed for %s: %v\n", str, err)
			continue
		}
		if err = v.Struct(pd); err != nil {
			fmt.Printf("Advanced validation failed for %s: %v\n\n", str, err)
			continue
		}
		fmt.Printf("Advanced validation passed for %s\n\n", str)
	}
}
