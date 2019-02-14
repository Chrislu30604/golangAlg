/*
Fibonacci Algorithm

F(n) = 1  ,             n = 1
	   1  ,             n = 2
	   F(n-1) + F(n-2), n > 2
*/
package main

import (
	"fmt"
	"math"
)

/*********************************
			  ALG 1
*********************************/
func fib(targetNumber uint) uint {
	n := targetNumber
	if n < 1 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// Time Complexity O(n)
// Extra Space O(n)

func multiply(F [][]uint, M [][]uint) {
	x := F[0][0]*M[0][0] + F[0][1]*M[1][0]
	y := F[0][0]*M[0][1] + F[0][1]*M[1][1]
	z := F[1][0]*M[0][0] + F[1][1]*M[1][0]
	w := F[1][0]*M[0][1] + F[1][1]*M[1][1]

	F[0][0] = x
	F[0][1] = y
	F[1][0] = z
	F[1][1] = w
}

/*********************************
			  ALG 2
*********************************/

// Optimizer
func power(F [][]uint, n uint) {
	if n == 0 || n == 1 {
		return
	}
	M := [][]uint{
		{1, 1},
		{1, 0},
	}

	power(F, n/2)
	multiply(F, F)

	if n%2 != 0 {
		multiply(F, M)
	}
}

func fib2(targetNumber uint) uint {
	F := [][]uint{
		{1, 1},
		{1, 0},
	}
	if targetNumber == 0 {
		return 0
	}
	power(F, targetNumber-1)
	return F[0][0]
}

// Time Complexity O(logN)
// Extra Space O(logN)

/*********************************
			  ALG 3
*********************************/
// use formulation Fn = {[(√5 + 1)/2] ^ n} / √5
func fib3(targetNumber uint) float64 {
	n := float64(targetNumber)
	phi := (1 + math.Sqrt(5)) / 2
	return math.Round(math.Pow(phi, n) / math.Sqrt(5))
}

func main() {
	var targetNumber uint = 20
	fmt.Println(fib(targetNumber))
	fmt.Println(fib2(targetNumber))
	fmt.Println(fib3(targetNumber))
}
