package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}

	oddEven := make(map[int][]int, 0)
	for _, x := range xs {
		if x % 2 == 0 {
			oddEven[0] = append(oddEven[0], x)
		} else {
			oddEven[1] = append(oddEven[1], x)
		}
	}

	fmt.Printf("%v\n", oddEven)
}