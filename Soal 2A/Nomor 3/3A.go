package main

import "fmt"

func main(){
	var r int
	var volumeBola, luasKulitBola float64

	fmt.Printf("Masukkan jari-jari: ")
	fmt.Scan(&r)

	volumeBola = (4.0 / 3.0) * 3.14 * float64(r) * float64(r) * float64(r)
	luasKulitBola = 4 * 3.14 * float64(r) * float64(r)

	fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f\n", r, volumeBola, luasKulitBola)
}