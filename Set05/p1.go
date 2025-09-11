package main

import (
	"errors"
	"fmt"
)

/*
   Problem 1: Implement a function to rotate
    a slice to the right by k steps.
*/

func rotateSlice(slice []int, num int) ([]int, error) {

	length := len(slice)

	if length < num {
		return nil, errors.New("invalid")
	}
	return append(slice[length-num:], slice[:length-num]...), nil
}

func P1() {

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	num := 5

	rotate, err := rotateSlice(slice, num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rotated slice ", rotate)
	}

}
