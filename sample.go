package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Pilihan 1")
		fmt.Println("2. Pilihan 2")
		fmt.Println("3. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		switch pilihan {
		case 1:
			fmt.Println("Anda memilih Pilihan 1")
		case 2:
			fmt.Println("Anda memilih Pilihan 2")
		case 3:
			fmt.Println("Keluar dari program")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
