package main

import "fmt"

func main() {
	var dias float64
	var diaria float64

	fmt.Print("Quantos dias voce trabalha por mês? ")
	fmt.Scan(&dias)
	fmt.Print("Quanto você recebe por dia? ")
	fmt.Scan(&diaria)

	salario := dias * diaria

	fmt.Print("O seu salário vai ser: ", salario)
}
