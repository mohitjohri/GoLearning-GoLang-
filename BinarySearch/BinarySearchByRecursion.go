package main

//BinarySearchByRecursion  Binary search by recursion function
func binarySearchByRecursion(slice []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	midIndex := int((lowIndex + highIndex) / 2)
	if slice[midIndex] > target {
		return binarySearchByRecursion(slice, target, lowIndex, midIndex)
	} else if slice[midIndex] < target {
		return binarySearchByRecursion(slice, target, midIndex+1, highIndex)
	} else {
		return midIndex
	}
}
