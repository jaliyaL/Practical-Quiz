package p1

import "fmt"

/*
Problem 1: Write a function to flatten a 2D slice of integers.
Input:

matrix := [][]int{
    {1, 2, 3},
    {4, 5},
    {6},
}

Output (flattened):

[]int{1, 2, 3, 4, 5, 6}

*/

func RunP1() {

	matrix := [][]int{{1, 2, 3}, {4, 5}, {6}}
	result := []int{}

	// version 01
	// for _, m := range matrix {
	// 	for _, v := range m {
	// 		result = append(result, v)
	// 	}
	// 	//fmt.Println(m)
	// }

	// version 2
	for _, m := range matrix {
		result = append(result, m...)
	}
	fmt.Println(result)
}
