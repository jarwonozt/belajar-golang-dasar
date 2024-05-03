package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// output negatif, karena melebihi kapasitas, sehingga kembali lagi ke belakang atau paling bawah
	fmt.Println(nilai16)

	var name = "Jarwonozt Aveiro"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
