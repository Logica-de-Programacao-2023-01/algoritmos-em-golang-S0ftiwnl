package main

import "fmt"

func main() {
	var kg float64

	fmt.Print("Quantos kg vc pesa? ")
	fmt.Scan(&kg)

	peso := kg * 2.20

	fmt.Print("Seu peso em libras é: ", peso)
}
