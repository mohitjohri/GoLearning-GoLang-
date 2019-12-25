package main

import (
	"fmt"
	"sort"

)

func main() {

	slice := []int{23, 56, 34, 45, 10, 5, 98, 58, 83, 113, 345, 21, 61, 75, 200}
	sort.Ints(slice) // this sort the slice of Int
	fmt.Println(slice)

	number := 58
	findByRecursion := binarySearchByRecursion(slice, number, 0, len(slice)-1)
	findByIteration := binarySearchByIteration(slice, number, 0, len(slice)-1)

	fmt.Printf("Position of number %d", number)
	fmt.Printf(" in the slice by recursion method is:%d\n", findByRecursion)
	fmt.Printf("Position of number %d", number)
	fmt.Printf(" in the slice by iteration method is:%d", findByIteration)
}
