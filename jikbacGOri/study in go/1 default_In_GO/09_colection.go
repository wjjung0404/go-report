package main

import (
	"fmt"
)

func main() {
	// array
	var array_collection = [3]int{1, 3, 3}
	fmt.Println(array_collection[1] == array_collection[2])
	fmt.Println(array_collection[1])
	for index, value := range array_collection {
		fmt.Println("index num :", index, "index value :", value)
	}
	// slice
	var slice_collection []int
	slice_collection = []int{1, 2, 3, 5, 6, 7}
	slice_collection[1] = 4
	fmt.Println("slice", slice_collection)
	fmt.Println("length : ", len(slice_collection), "cap : ", cap(slice_collection))
	for index, value := range slice_collection {
		fmt.Println("slice index : ", index, "slice value : ", value)
	}
	// cap, length can control
	var slice_control = make([]int, 5, 10)
	fmt.Println("len : ", len(slice_control), "cap : ", cap(slice_control))

	// sub slice
	var slice_sub = slice_collection[2:5]
	fmt.Println(slice_sub)

	//append , copy
	a := make([]int, 0, 10)
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println("a = append")
	fmt.Println(a)

	b := []int{11, 12, 123, 14, 15}
	a = append(a, b...)
	fmt.Println("b append the a", a)

	copy(a, b)
	fmt.Println("copy b to a", a)

	// MAP
	// map[key-type]value-type
	var defaultMAP map[int]string
	defaultMAP[111] = "aaa"
	defaultMAP[112] = "aab"
	defaultMAP[101] = "a^a"
	// or
	defaultMAP_make := make(map[int]string)
	defaultMAP_make[1] = "aaa"
	defaultMAP_make[1] = "aab"
	defaultMAP_make[1] = "a^a"
	// use leterer
	defaultMAP_literal := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	fmt.Println(defaultMAP_literal)

	// read the map-type variable value
	fmt.Println("[111]의 값, [101]의 값, make[1]의 값, make[4](없는 값)의 값")
	fmt.Println(defaultMAP[111], defaultMAP[101], defaultMAP_make[1], defaultMAP_make[4])

	//list delete
	delete(defaultMAP, 111)
	fmt.Println(defaultMAP[111]) // nil

	// map key check
	// GO list return a value 2 type
	val, exists := defaultMAP[112]
	if !exists {
		fmt.Println(exists, "NO have key the list")
	}
	fmt.Println("val = ", val, "exists = ", exists)

	// enum a list use for
	for key, val := range defaultMAP {
		fmt.Println(key, val)
	}
}
