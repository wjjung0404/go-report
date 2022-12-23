package main

import "fmt"

func main() {
	// func are devide a function
	// func funcname(argument) { }
	// use funcname(parameter)
	fmt.Println("I'am fmt libray's Function")
	fmt.Println(something("lol"))
	a, b, res := sum(10, 20)
	fmt.Println(a, b, res)
}

// func funcname( argument ) return type
// argument, return type can remove
func something(think string) string {
	input := "your think : " + think
	// return are go to called
	return input
}

// and many use argument and return type
// we use array is next time Collection
func sum(a, b int32) (int32, int32, int32) {
	res := a + b
	return res, a, b
}

// or should this
func sum_many(val ...int32) int32 {
	var res int32 = 0
	for index, oper := range val {
		res += oper
		fmt.Println(index)
	}
	return res
}
