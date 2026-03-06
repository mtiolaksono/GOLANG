package main

import "fmt"

func main() {
	// Cara Pertama membuat variabel
	var name1 string

	name1 = "muhammad tio lakasono"
	fmt.Println(name1)

	// Cara kedua membuat variabel
	name2 := "muhammad tio laksono"

	fmt.Println(name2)

	// Cara ketiga membuat varibael
	var (
		firstname  = "Muhammad"
		middlename = "Tio"
		lastname   = "Laksono"
	)

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
	fmt.Println(firstname + " " + middlename + " " + lastname)
}
