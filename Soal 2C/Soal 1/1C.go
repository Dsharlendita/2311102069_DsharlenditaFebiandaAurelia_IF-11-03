package main

import "fmt"

func main() {
	var beratGram, beratKg, sisaGram, biayaKg, biayaSisa, totalBiaya  int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratGram)

	beratKg = beratGram / 1000
	sisaGram = beratGram % 1000
	biayaKg = beratKg * 10000

	if beratKg >= 10 {
		biayaSisa = 0 
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Printf("Berat parsel (gram): %d\n", beratGram)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
