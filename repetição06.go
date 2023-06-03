package main

import "fmt"

func main() {
	var y int

	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)

	for x := 0; x <= 10; x++ {
		fmt.Print(" ", x*y)
	}
}
