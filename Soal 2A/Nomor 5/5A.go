package main

import "fmt"

func main(){
	var angka1, angka2, angka3, angka4, angka5 int
	var huruf1, huruf2, huruf3 byte

	fmt.Scanln(&angka1, &angka2, &angka3, &angka4, &angka5)
	fmt.Scanf("%c%c%c", &huruf1, &huruf2, &huruf3)

	fmt.Printf("%c%c%c%c%c\n", angka1, angka2, angka3, angka4, angka5)
	fmt.Printf("%c%c%c\n", huruf1+1, huruf2+1, huruf3+1)
}