package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Print(" FizzBuzz")
		} else if x%3 == 0 {
			fmt.Print(" Fizz")
		} else if x%5 == 0 {
			fmt.Print(" Buzz")
		} else {
			fmt.Print(" ", x)
		}
	}
}
