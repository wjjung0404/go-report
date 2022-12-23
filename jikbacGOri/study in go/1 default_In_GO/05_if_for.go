package main

import "fmt"

func main() {
	// if expect a value are true
	if 10 > 11 {
		fmt.Println("10 > 6 is true")
	} else if 10 == 11 { // 10 > 11 is not true, next compare is here
		fmt.Println("10 == 11 is true")
	} else { // 1, 2 compare not true, go to here
		fmt.Println("10 >= 11 is false")
	}
	// else if, else can not use

	// switch is expectally are many used convision a value
	// failThrough down to next case at true case (=OR)
	val := 10
	switch val {
	case 11:
		fmt.Println("val = 11")
	case 10:
		fmt.Println("val = 10")
		fallthrough
	case 9:
		fmt.Println("val = 9")
	default:
		fmt.Println("default is every conpare when false, go to here")
	}
	// go can use Express in if,switch
	if i := 5; i > 10 {
		fmt.Println("i value up as 10")
	}
	switch i := 5; i > 10 {
	case true:
		fmt.Println("true")
	default:
		fmt.Println("false ")
	}

	// for
	// for are loop block
	for i := 0; i <= 5; i++ {
		fmt.Println("index i : ", i)
	}
	my_for1 := 0
	for my_for1 < 10 {
		my_for1 = my_for1 + 1
		fmt.Println("my_for1 : ", my_for1)
	}
	// infinity loop
	// waring : it's very danger because infinity loop, not next line,
	// and memory over
	loop_i := 10
	for {
		loop_i++ // loop_i + loop_i + 1
		if loop_i > 15 {
			break
		}
	}
}
