package main

import "fmt"

func main() {
	for x := 10; x <= 10 && x > 0; x-- {
		fmt.Print(" ", x)
	}
}
