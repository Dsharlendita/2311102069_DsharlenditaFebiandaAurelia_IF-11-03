package main

import "fmt"

func main() {
	var kantongTerpalKanan, kantongTerpalKiri, selisih float64

	for {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantongTerpalKanan, &kantongTerpalKiri)

		if kantongTerpalKanan < 0 || kantongTerpalKiri < 0 || kantongTerpalKanan+kantongTerpalKiri > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih = kantongTerpalKanan - kantongTerpalKiri
		if selisih < 0 {
			selisih = -selisih 
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
