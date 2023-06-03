package main

import "fmt"

func main() {
	var idade float64

	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)

	fmt.Print("A sua idade em dias é: ", idade*365)
}
