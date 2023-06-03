package main

import "fmt"

func main() {
	for x := 0; x >= 0 && x <= 30; x += 3 {
		fmt.Print(" ", x)
	}
}
