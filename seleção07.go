package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)

	mais := salario * 0.05
	menos := salario * 0.1

	if salario < 1000 {
		fmt.Print("Você obteve um aumento de 10% e seu novo salário é: ", salario+menos)
	} else {
		fmt.Print("Você obteve um aumento de 5% e seu novo salário é: ", salario+mais)
	}
}
