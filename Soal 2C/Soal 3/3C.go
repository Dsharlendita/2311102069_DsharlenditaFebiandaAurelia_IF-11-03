package main

import "fmt"

func main() {
	var b int

	fmt.Print("Masukkan bilangan bulat positif (lebih dari 1): ")
	fmt.Scan(&b)

	fmt.Printf("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	prima := true
	if b <= 1 {
		prima = false
	} else {
		for i := 2; i*i <= b; i++ {
			if b%i == 0 {
				prima = false
				break
			}
		}
	}
	fmt.Printf("Prima: %v\n", prima)
}
