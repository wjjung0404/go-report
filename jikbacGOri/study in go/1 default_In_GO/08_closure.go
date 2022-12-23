package main

func main() {
	var ai = nextValue()
	println(ai())
	println(ai())
	println(ai())
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
