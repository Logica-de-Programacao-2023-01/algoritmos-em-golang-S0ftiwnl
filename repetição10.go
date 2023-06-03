package main

import "fmt"

func main() {
	var (
		numero, maior    int
		primeiraIteracao = true
	)

	for {
		fmt.Print("Digite um número inteiro (digite 0 para sair): ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		if primeiraIteracao || numero > maior {
			maior = numero
			primeiraIteracao = false
		}
	}

	if primeiraIteracao {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		fmt.Printf("O maior número digitado é: %d\n", maior)
	}
}
