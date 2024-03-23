package main

import (
	"fmt"
)

func Variable() {
	var nums int = 32// declared a variable nums with type int
	var floats float64 = 3.3368
	var isNum bool = true
	var hello string = "Hello"
	fmt.Printf("%v, %.2f %v %q \n", nums, floats, isNum, hello)

	// convert type
	a := 90.99 // 90.99
	b := int(a) // 90 -> this value had truncated
	fmt.Println(a, b)

	// constant
	const aString string = "Hellllo"
	const isHatttt bool = true
	const noNo int = 90

	// fmt.Printf -> formats a string and print it to standard out
	// fmt.Sprintf -> formats a string and return the string
}