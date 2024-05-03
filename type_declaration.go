package main

import "fmt"

func main() {
	type NoKtp string

	var ktpNozt NoKtp = "12345556953"

	var contoh string = "22222121212"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpNozt)
	fmt.Println(contohKtp)
}
