package main

import "fmt"

func main() {
	var dia int

	fmt.Print("Digite um valor entre 1 e 7: ")
	fmt.Scan(&dia)

	if dia == 1 {
		fmt.Print("1 é domingo")
	} else if dia == 2 {
		fmt.Print("2 é segunda")
	} else if dia == 3 {
		fmt.Print("3 é terça")
	} else if dia == 4 {
		fmt.Print("4 é quarta")
	} else if dia == 5 {
		fmt.Print("5 é quinta")
	} else if dia == 6 {
		fmt.Print("6 é sexta")
	} else if dia == 7 {
		fmt.Print("7 é sábado")
	} else {
		fmt.Print("Valor inválido")
	}
}
