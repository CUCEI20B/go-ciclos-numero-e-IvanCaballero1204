package main

import "fmt"

func main() {
	var n int
	var e float64 = 0.0

	factorial := 1.0

	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		for j := i; j > 1; j-- {
			factorial *= float64(j)
		}
		e += 1.0 / float64(factorial)
		factorial = 1
	}
	println(e)
}
