package main

import (
	"fmt"
	"reflect"
)

func main() {
	// primitive type: boolean
	var isGoAwesome bool
	isGoAwesome = false // just kidding!
	isGoAwesome = true
	fmt.Println(isGoAwesome)

	// primitive type: integer
	var integer int
	integer = -10
	integer = 10
	fmt.Println(integer)

	var followerCount uint32
	followerCount = 10000
	// overflow error
	// followerCount = -1000
	fmt.Println(followerCount)

	// primitive type: float
	var balance float32
	balance = 10000.01
	fmt.Println(balance)

	// primitive type: string
	var text string
	text = "Go is fun."
	fmt.Println(text)

	// an easier way to declare variables
	// import "reflect"
	isGoAwesomeNew := true
	fmt.Println(reflect.TypeOf(isGoAwesomeNew))
	integerNew := 100
	fmt.Println(reflect.TypeOf(integerNew))
	floatNew := 100.1
	fmt.Println(reflect.TypeOf(floatNew))
	textNew := "Go is fun"
	fmt.Println(reflect.TypeOf(textNew))
}
