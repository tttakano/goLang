package main

import "fmt"

func main(){
	xs := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, x := range xs {
		sum += x
	}

	fmt.Printf("%d\n", sum)
}