package main

import "fmt"

func main() {

	ints := []int{1, 6, 5, 4, 3, 2, 1}

	for i := 1; i < len(ints); i++ {

		if ints[i-1] > ints[i] {
			ints[i-1], ints[i] = ints[i], ints[i-1]
		}

	}

	fmt.Println(ints)

}
