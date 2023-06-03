package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)
	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Print("A multiplicação entre x e y é: ", x*y)
	} else {
		fmt.Print("A soma entre x e y é: ", x+y)
	}
}
