package main

import "fmt"

func main() {
	// Golang's variable type have a primitive type
	// 1. bool type
	var is_bool bool = true
	fmt.Println(is_bool)
	// 2. string
	var str = "Hello world" // string type are immutable
	fmt.Println(str)
	// 3. integer
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	var integer_type1 int32 = 32
	fmt.Println(integer_type1)
	// 4. float
	// float32 float64 complex64 complex128
	var float_type1 float32 = 3.14
	fmt.Println(float_type1)
	// 5. other type
	// byte : use byte code, same unit8
	// rune : used in unicode codepoint, same int32

	// String type are specially
	// string can make a use ``(back quote) OR "":
	// use "" make a Interpreted String Literal
	// it is used escape(\n, \t ....)
	var preted_str = "A\tB\tC\nHello"
	fmt.Println(preted_str)
	// use ``(back quote) make a Raw String Literal
	// not used escape, just take raw string
	var raw_str = `A\tB\tC\nHello`
	fmt.Println(raw_str)

	// Type Conversion
	// one of convers A to B, use 'DataType(variable[or value])' syntax
	int32_a := 11
	var int32_float32 = float32(int32_a) + 0.1
	fmt.Println(int32_float32)
}
