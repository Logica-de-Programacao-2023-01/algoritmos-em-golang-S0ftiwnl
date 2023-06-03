package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Print("Você é considerado mirim")
	} else if idade >= 10 && idade <= 13 {
		fmt.Print("Você9-1" +
			" é considerado infantil")
	} else if idade >= 14 && idade <= 17 {
		fmt.Print("Você é considerado juvenil")
	} else {
		fmt.Print("Você é considerado adulto")
	}
}
