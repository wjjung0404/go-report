package main

import "fmt"

func main() {
	// operater are can operation
	// +, -, *, /, % are math operator
	// "/" are value of division, "%" are value of division remainder
	var math = (3 + 5*6/7)
	fmt.Println(math)
	// <, <=, >, >=, ||(OR), &&(AND), ! are logical(true, false) operator
	var logi = 5 < 6 && 7 > 8
	fmt.Println(logi)
	// <<, >>, &, | are bitwise operator (0100 => 1000)
	var bit = (2 << 1) | 3
	fmt.Println(bit)
	// &a, *a are pointer operator
	var my_val = 10
	var my_val_memory = &my_val
	fmt.Println(my_val_memory, *my_val_memory)
	// operator have priority
	/*
		priority list
		1.	*,/,%, <<, >>, &, &^
		2.	+,-, \, ^\
		3.  ==, !=, <, <=, >, >=
		4.	&&
		5.	||
	*/
	// left to right, top to down
}
