package main

import "fmt"

/*
Problem 1: Write a function to remove duplicates from a slice of integers.
*/
func removeDuplicates(s []int) []int {

	check := make(map[int]bool)
	cleanSlice := make([]int, 0)
	duplicateSlice := make([]int, 0)

	for _, val := range s {
		if _, exists := check[val]; exists {
			duplicateSlice = append(duplicateSlice, val)
		} else {
			check[val] = true
			cleanSlice = append(cleanSlice, val)
		}
	}
	fmt.Println(check)
	fmt.Println("duplicate values", duplicateSlice)
	return cleanSlice
}

func P1() {

	oSlice := []int{1, 5, 6, 1, 4, 2, 8, 9, 2, 5}
	nSlice := removeDuplicates(oSlice)
	fmt.Println("original", oSlice)
	fmt.Println("clean", nSlice)
}
