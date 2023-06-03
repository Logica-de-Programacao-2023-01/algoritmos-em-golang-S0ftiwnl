package main

import "fmt"

func main() {
	for x := 0; x >= 0 && x <= 20; x += 2 {
		fmt.Print(" ", x)
	}
}
