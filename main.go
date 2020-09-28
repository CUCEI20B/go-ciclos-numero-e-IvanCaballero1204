package main

import "fmt"

func main() {
	var limite int
	var euler float64
	var factorial float64

	fmt.Scan(&limite)

	for i := 0; i <= limite; i++ {
		factorial = 1
		for j := i; j > 1; j-- {
			factorial = factorial * float64(j)
		}
		
		euler += (1 / factorial)
	}
	fmt.Println(euler)
}
