package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)
	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)
	fmt.Print("Qual o valor de z? ")
	fmt.Scan(&z)

	var mediaponderada float64
	mediaponderada = (x*2 + y*3 + z*5) / 10
	fmt.Print("A média ponderada dos valores inseridos é:", mediaponderada)
}
