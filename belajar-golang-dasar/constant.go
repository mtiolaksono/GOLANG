package main

import "fmt"

func main() {
	// Cara pertama membuat constant
	const firstname1 string = "Muhammad Tio"
	const lastname1 = "Laksono"

	fmt.Println(firstname1)
	fmt.Println(lastname1)

	// Cara kedua membuat constant

	const (
		firstname2 = "Muhammad Tio"
		lastname2  = "Lakasono"
	)

	fmt.Println(firstname2)
	fmt.Println(lastname2)
}
