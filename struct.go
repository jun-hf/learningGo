package main

import "fmt"

func EmbedeStruct() {
	type Wheel struct {
		nums int
		grade int
	}

	type Car struct {
		Wheel
		grade int
	}

	fmt.Println(Car{}.Wheel.grade)
}