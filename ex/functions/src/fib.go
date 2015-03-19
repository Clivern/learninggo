package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) |\longremark{At \citem{} we create an \key{array} to hold the integers up to the value given in the function call.}|
	x[0], x[1] = 1, 1 |\longremark{At \citem{} we have the start of the Fibonacci calculation.}|
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2] |\longremark{$x_n = x_{n-1} + x_{n-2}$.}|
	}
	return x |\longremark{At \citem{} we return the \emph{entire} array.}|
}

func main() {
	for _, term := range fibonacci(10) { |\longremark{And at \citem{} we use the \key{range} keyword ito  ``walk'' the numbers returned by the Fibonacci function. Here up to 10. Finally, we print the numbers.}|
		fmt.Printf("%v ", term)
	}
}
