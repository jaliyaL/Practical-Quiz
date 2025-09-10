package p2

import "fmt"

/*
Problem 2: Implement a concurrent Fibonacci generator using channels.
*/
func fibo(n int) int {
	if n <= 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

func RunP2() {
	// normal finonachi  without channels
	// var r int
	for i := 0; i < 10; i++ {
		r := fibo(i)
		fmt.Println(r)
	}

}
