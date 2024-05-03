package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jarwonozt"
	names[1] = "Aveiro"
	names[2] = "Bailey"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Array langsung
	var values = [...]int{
		90,
		80,
		70,
		170,
		270,
		370,
	}

	fmt.Println(values)

}
