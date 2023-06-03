package main

import "fmt"

func main() {
	var produto float64

	fmt.Print("Qual o valor do seu produto? ")
	fmt.Scan(&produto)

	desconto := produto * 0.1

	fmt.Print("Sua compra com desconto será: ", produto-desconto)
}
