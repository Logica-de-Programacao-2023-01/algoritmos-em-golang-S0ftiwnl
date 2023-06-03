package main

import "fmt"

func main() {
	var x int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)

	if (x%3 == 0 && x%5 == 0) {
		fmt.Print("X eh divisível por 3 e 5 ao mesmo tempo")
	} else {
		fmt.Print("X não é divisível, por 3 e 5 ao mesmo tempo")
	}
}
