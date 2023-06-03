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

	if x < y && x < z {
		fmt.Print("X é o menor valor entre os 3")
	} else if y < x && y < z {
		fmt.Print("Y é o menor valor entre os 3")
	} else {
		fmt.Print("Z é o menor valor entre os 3")
	}
}
