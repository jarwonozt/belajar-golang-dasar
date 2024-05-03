package main

import "fmt"

func main() {

	/*
		kata kunci var tidak lah wajib
		asalakan kita langsung menginisialisasikan datanya
		agat tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci :=
		saat menginisialisasikan data pada variable tersebut
	*/
	// var name string
	title := "Perjalanan baru bersama golang"

	fmt.Println(title)

	name := "Jarwonozt"
	fmt.Println(name)

	name = "Jarwonozt Aveiro"
	fmt.Println(name)

	// Multiple variable
	var (
		firstName = "Jarwonozt"
		lastName  = "Aveiro"
	)

	fmt.Print(firstName)
	fmt.Print(lastName)

}
