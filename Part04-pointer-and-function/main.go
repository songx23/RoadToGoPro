package main

import "fmt"

func main() {
	pointer := "I'm a string, not a pointer"
	var pointerType *string
	pointerType = &pointer
	// declare directly
	p := &pointer
	fmt.Printf("String: %s\nPointerType: %v\nPointer: %v\n", pointer, pointerType, p)
	fmt.Printf("Dereference pointer: %v\n", *p)

	// Uncomment below to see the famous nil pointer exception
	// var nilPointer *int
	// fmt.Printf("nil pointer: %v\n", *nilPointer)
	// output:
	//    panic: runtime error: invalid memory address or nil pointer dereference
	*p = "Pretending to be a pointer now."
	fmt.Printf("New value set via pointer: %v\n", *p)
	fmt.Printf("Variable value changed: %s\n", pointer)
}
