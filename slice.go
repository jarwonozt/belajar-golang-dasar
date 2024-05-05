package main

import "fmt"

func main() {
	names := [...]string{"Jarwonozt", "Dos", "Santos", "Aveiro", "Pamungkas", "Ferdinan"}

	// mengambil data dari index ke-4 sampai index ke-6
	slice1 := names[4:6]

	slice2 := names[:4]
	slice3 := names[3:]
	slice4 := names[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Juma'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // [Senin, Selasa, Rabu, Kamis, Juma'at, Sabtu Baru, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // [Ups, Minggu Baru, Libur Baru]
	fmt.Println(days)      // [Senin, Selasa, Rabu, Kamis, Juma'at, Sabtu Baru, Minggu Baru]

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Jarwonozt"
	newSlice[1] = "Aveiro"
	// newSlice[2] = "Santos" //error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// append
	newSlice2 := append(newSlice, "Santos")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

}
