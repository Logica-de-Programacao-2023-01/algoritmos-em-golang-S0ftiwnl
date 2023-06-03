package main

import "fmt"

func main() {
	for x := 1; x >= 1 && x <= 19; x += 2 {
		fmt.Print(" ", x)
	}
}
