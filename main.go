package main

import (
	"fmt"
	"os"
	"strconv"
)

/*func main() {
	size := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cumSum := 0

	//only one or the other of these for loops
	//these are just to see different ways numbers from 1-10 can be summed

	for i := 0; i <= len(size); i++ {
		cumSum += size[i]
	}
	fmt.Printf("Sum: %d\n", cumSum)

	for _, value := range size {
		cumSum += value
	}
	fmt.Printf("Sum: %d\n", cumSum)
}

func sum() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)
} */

/*func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		if i != 10 {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Print(" = ", sum)
	fmt.Println()
} */

func main() {
	var sum int

	if len(os.Args) != 3 {
		fmt.Println("Please enter a minimum and maximum number to add.")
		return
	}

	min := os.Args[1]
	max := os.Args[2]

	if minInt, err1 := strconv.Atoi(min); err1 != nil {
		fmt.Printf("%s is not a number. \n", min)
		return
	} else if maxInt, err2 := strconv.Atoi(max); err2 != nil {
		fmt.Printf("%s is not a number. \n", max)
		return
	} else {
		for i := minInt; i <= maxInt; i++ {
			sum += i
			if i != maxInt {
				fmt.Printf("%d + ", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
	}
	fmt.Print(" = ", sum)
	fmt.Println()

}
