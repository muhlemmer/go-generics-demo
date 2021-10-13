package main

import "fmt"

// convertible integer types.
type convertible interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func convert[A, B convertible](a A) B {
	return B(a)
}

func main() {
	var a int8 = 127

	b := convert[int8, uint32](a)

	fmt.Printf("Converted %T to %T with value of %d\n", a, b, b)
}
