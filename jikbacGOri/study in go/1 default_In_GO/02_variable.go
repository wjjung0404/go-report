package main

import "fmt"

func main() {
	// You want a comment in your code
	// use commently "//"
	// fmt.Print("Hello world");

	// GO lang's variable declear can write a var first
	// and next wriet variable name, last declaer a this type
	var integer_a_variable int
	// variable value initial setting can decleared variable on next
	integer_a_variable = 8
	fmt.Print(integer_a_variable)

	// or can variable declear together a initialize
	var float_b_barialbe = 11.0
	fmt.Print(float_b_barialbe)

	// if you declear a variable here, variable should be used
	// go lang are not allow variable that not used
	// if you erase a commentary('//') at 'var write_a_var = 10'(next line),
	// it's occurs a error because not used in program
	// var write_a_var = 10
	// so when you declear variable in project, use delceared variable is important

	// go lang can many declaer variable in 1 line together
	var a, b, c int = 1, 2, 3
	fmt.Print(a, b, c)

	// variable decleared zero value when it's not initalize value
	var zero int // it's have a 0
	fmt.Print(zero)

	// another one use ':=' when just in function
	use_short_var := 10
	fmt.Print(use_short_var)

	// const variable is immutable so can't change value
	const immutable_variable = 10
	// if you need many immutable value, can use it
	const (
		const_1000 = 5
		const_500  = 4
		const_100  = 2
	)
	// or need a ordered value, should use iota
	const (
		flag1 = iota // 0
		flag2        // 1
		flat3        // 2 ...
	)

	// when you declear a variable, can not use 25 keyword your variable name
	// it is used in this languege
	/*
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	*/
}
