package main

import "fmt"

func main() {
	var peso float64
	var altura float64

	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)
	fmt.Print("Qual a sua altura? ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Print("Seu imc é: ", imc)

	if imc <= 18.5 {
		fmt.Print("\nVocê está abaixo do peso")
	} else if imc > 18.5 && imc <= 24.9 {
		fmt.Print("\nVocê está com o peso normal")
	} else {
		fmt.Print("\nVoce está com sobrepeso")
	}
}
