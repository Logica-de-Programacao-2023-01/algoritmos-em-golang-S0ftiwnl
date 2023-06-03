package main

import "fmt"

func main() {
	for x := 1; x >= 0 && x <= 10; x += 2 {
		fmt.Print(" ", x)
	}
}
