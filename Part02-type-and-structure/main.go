package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ***************
	// primitive types
	// ***************
	fmt.Println("Primitive types:")
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

	// ***************
	// composite types
	// ***************
	fmt.Println("Composite types:")
	// composite type: array
	var shoppingList [3]string
	shoppingList[0] = "Milk"
	shoppingList[1] = "Avocado" //smash it
	shoppingList[2] = "Beer"
	fmt.Println(shoppingList)
	// error if you want to stuff more elements than the array can handle
	//   * uncomment the next line to see the error
	// shoppingList[3] = "out of bounds"

	// or declare arrays in a simpler way
	simpleShoppingList := [3]string{"Milk", "Avocado", "Beer"}
	fmt.Println(simpleShoppingList)

	// or make compiler do even more work, in case you don't want to count the size
	lazyShoppingList := [...]string{"Milk", "Avocado", "Beer"}
	fmt.Println(lazyShoppingList)

	// to get avocado from my shopping list, I need to find its index
	avocado := shoppingList[1]
	// fmt package can do better than this but let's stick with Println for now
	fmt.Println("Smash ", avocado, "\n")

	// composite type: slice
	// slice can stuff as many kinds of bread as you can imagine
	bread := []string{"ficelle", "la baguette", "brioche"}
	fmt.Println(bread)

	// another way to declare slice
	makeBread := make([]string, 3)
	makeBread[0] = "ficelle"
	makeBread[1] = "la baguette"
	makeBread[2] = "brioche"
	fmt.Println(makeBread)

	// creating slice from array/slice
	breadArr := [3]string{"ficelle", "la baguette", "brioche"}
	breadSlice := breadArr[0:3]
	fmt.Println(breadSlice)

	// append new elements to a slice
	makeBread = append(makeBread, "new bread")
	fmt.Println(makeBread)

	// slicing from existing array
	breadSliceExample1 := breadArr[0:2]
	fmt.Println(breadSliceExample1) // [ficelle la baguette]
	breadSliceExample2 := breadArr[1:3]
	fmt.Println(breadSliceExample2) // [la baguette brioche]]

	// composite type: map
	postcodes := make(map[int]string)
	postcodes[3000] = "Melbourne"
	postcodes[2000] = "Sydney"
	postcodes[6000] = "Perth"
	fmt.Println(postcodes)

	// declare map inline with values
	postcodesInline := map[int]string{
		3000: "Melbourne",
		2000: "Sydney",
		6000: "Perth",
	}
	fmt.Println(postcodesInline)
}
