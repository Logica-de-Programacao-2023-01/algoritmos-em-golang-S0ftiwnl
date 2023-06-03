package main

import "fmt"

func main() {
	var peso float64
	var altura float64

	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)
	fmt.Print("Qual a sua altura? ")
	fmt.Scan(&altura)

	fmt.Print("O seu imc é: ", peso/(altura*altura))
}
