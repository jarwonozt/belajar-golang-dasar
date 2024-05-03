package main

import (
	"fmt"
)

func main() {
	/*
		Constant
		Variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	*/

	const firstName string = "Jarwonozt"
	const lastName = "Aveiro"

	// error
	// firstName = "Mesut"
	// lastName = "Ozil"

	fmt.Println(firstName)
	fmt.Println(lastName)

}
