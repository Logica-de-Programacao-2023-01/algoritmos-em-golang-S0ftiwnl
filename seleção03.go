package main

import "fmt"

func main() {
	var x int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Print("X é par")
	} else {
		fmt.Print("X é ímpar")
	}
}
