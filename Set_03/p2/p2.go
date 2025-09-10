package p2

import "fmt"

/*
Problem 2: Implement a concurrent Fibonacci generator using channels.
*/

// without channels
// func fibo(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	return fibo(n-1) + fibo(n-2)
// }

// func RunP2() {
// 	// normal finonachi  without channels
// 	// var r int
// 	for i := 0; i < 10; i++ {
// 		r := fibo(i)
// 		fmt.Println(r)
// 	}
// }

// with channels
func fibonachi(n int, ch chan int) {

	x, y := 0, 1
	for i := 0; i <= n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)

}

func RunP2() {

	ch := make(chan int, 10)
	go fibonachi(10, ch)

	for c := range ch {
		fmt.Println(c)
	}
}
