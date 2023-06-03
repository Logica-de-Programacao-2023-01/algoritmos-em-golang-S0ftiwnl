package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)
	fmt.Print("Digite o valor de y: ")
	fmt.Scan(&y)
	fmt.Print("Digite o valor de z: ")
	fmt.Scan(&z)

	if x < y && x < z {
		fmt.Print("Os números em ordem crescente: ", x, y, z)
	} else if x < y && y < z {
		fmt.Print("Os números em ordem crescente: ", x, z, y)
	} else if y < x && x < z {
		fmt.Print("Os números em ordem crescente: ", y, x, z)
	} else if y < z && z < x {
		fmt.Print("Os números em ordem crescente: ", y, z, x)
	} else if z < x && x < y {
		fmt.Print("Os números em ordem crescente: ", z, x, y)
	} else {
		fmt.Print("Os números em ordem crescente: ", z, y, x)
	}

}
