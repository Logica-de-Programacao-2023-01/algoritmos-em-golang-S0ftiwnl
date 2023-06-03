package main

import "fmt"

func main() {
	var x int

	fmt.Printf("Qual o valor de x? ")
	fmt.Scan(&x)

	fmt.Print("O dobro, triplo e quádruplo de x é ", x*2, x*3, x*4)
}
