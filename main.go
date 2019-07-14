package main

import "fmt"

func main() {
	// array size is part of its type
	// initialization with default values
	var a1 [3]int // [0,0,0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool // [false,false,false,false]
	fmt.Println("a2", a2)

	// specification of array size at declaration
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	// verification at compilation or run-time
	// invalid array index 4 (out of bounds for 3-element array)
	// a3[idx] = 12
}
