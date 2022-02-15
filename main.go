package main

import "fmt"

func main() {
	size := 10

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("Sum: %d\n", i+j)
			fmt.Println()
		}
	}
}
