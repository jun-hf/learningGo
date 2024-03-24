package main

import "fmt"

func pointer() {
	v := 90
	p := &v
	*p = 89

	fmt.Printf("v: %v", v)
}