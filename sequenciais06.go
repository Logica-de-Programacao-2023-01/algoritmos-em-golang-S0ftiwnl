package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)

	var aumento float64
	aumento = salario * 0.15

	fmt.Print("O seu salario com um aumento de 15% é: ", salario+aumento)
}
