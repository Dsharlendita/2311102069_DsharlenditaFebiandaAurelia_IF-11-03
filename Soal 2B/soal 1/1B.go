package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var status bool = true

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			status = false 
		}
	}
	if status {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
