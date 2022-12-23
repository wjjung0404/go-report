package main

import "fmt"

func main() {
	// anonymous function is not have function name but can used
	// anonymous function can init to variable
	compare := func(a int, b int) bool {
		return a > b
	}
	fmt.Println(compare(2, 3))
	// first-class function
	// first-class function same primitive type( int, int32, float32...)
	// so Go lang's function can use a function's parameter or return value
	r1 := sum(func(x, y int) int { return x - y }, 10, 20)
	fmt.Println(r1)
}

// parameter f1 is function's function parameter
func sum(f1 func(int, int) int, v1, v2 int) int {
	res := v1 + v2 + f1(v1, v2)
	return res
}
