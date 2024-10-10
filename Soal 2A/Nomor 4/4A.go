package main

import "fmt"

func main(){
	var celcius, reamur, fahrenheit, kelvin float64
	
	fmt.Printf("Masukkan temperatur Celcius: ")
	fmt.Scan(&celcius)

	fahrenheit = celcius * 9.0 / 5.0 + 32
	reamur = celcius * 4.0 / 5.0
	kelvin = 5.0 / 9.0 * (fahrenheit + 459.67)

	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Kelvin: ", int(kelvin))
}