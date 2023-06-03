package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)
	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)
	fmt.Print("Qual o valor de z? ")
	fmt.Scan(&z)

	fmt.Print("A soma dos valores é: ", x+y+z)
}
