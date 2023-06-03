package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)
	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)

	if x > y {
		fmt.Print("X é maior que y")
	} else {
		fmt.Print("Y é maior que x")
	}
}
