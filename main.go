package main

import "fmt"

func main() {
	size := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cumSum := 0

	//only one or the other of these for loops
	//these are just to see two different ways it can be done

	for i := 1; i <= len(size); i += 2 {
		cumSum += size[i] + size[i-1]
	}
	fmt.Printf("Sum: %d\n", cumSum)

	for _, value := range size {
		cumSum += value
	}
	fmt.Printf("Sum: %d\n", cumSum)
}
