package main

import "fmt"

func main() {
	var tahun int
	fmt.Printf("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Printf("Tahun: %d\nKabisat: true\n", tahun)
	} else {
		fmt.Printf("Tahun: %d\nKabisat: false\n", tahun)
	}
}