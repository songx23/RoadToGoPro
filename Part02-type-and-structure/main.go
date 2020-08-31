package main

import (
	"fmt"
	"reflect"
	"strconv"
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
	fmt.Println("Smash ", avocado)

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

	// access value in map using key
	melbourne := postcodes[3000]
	fmt.Println(melbourne) // output: Melbourne

	// access non-existing key
	darwin := postcodes[8000]
	fmt.Println("darwin is empty string? ", darwin == "") // darwin: empty string

	// check key existed or not
	darwin, ok := postcodes[8000]
	fmt.Println("is key present in map? ", ok)            // output: false
	fmt.Println("darwin is empty string? ", darwin == "") // darwin: empty string

	// add key value pair
	postcodes[4000] = "Brisbane"
	fmt.Println(postcodes) // should see a new pair in map

	// delete key value pair
	delete(postcodes, 4000)
	fmt.Println(postcodes) // should see 4000:Brisbane is removed from map

	// Composite type: struct
	// declare an account
	var account1 Account
	account1.BSB = 123456
	account1.AccountNumber = 12345678
	account1.AccountType = "Joint Account"
	account1.Owners = []Owner{
		{
			FirstName:   "Jon",
			LastName:    "Snow",
			PhoneNumber: "0400100200",
			Email:       "jon.snow@got.com",
		},
		{
			FirstName:   "Daenerys",
			LastName:    "Stormborn",
			PhoneNumber: "0400300400",
			Email:       "mother.of.dragons@got.com",
		},
	}
	fmt.Printf("%+v\n", account1)

	// another way to declare an account inline
	account2 := Account{
		BSB:           123456,
		AccountNumber: 12345678,
		AccountType:   "Joint Account",
		Owners: []Owner{
			{
				FirstName:   "Jon",
				LastName:    "Snow",
				PhoneNumber: "0400100200",
				Email:       "jon.snow@got.com",
			},
			{
				FirstName:   "Daenerys",
				LastName:    "Stormborn",
				PhoneNumber: "0400300400",
				Email:       "mother.of.dragons@got.com",
			},
		},
	}
	fmt.Printf("%+v\n", account2)

	// Composite type: interface
	// declare a variable with interface type
	var poodle Animal
	poodle = Dog{Name: "Bolt"}
	fmt.Println(poodle.Eat())

	// ***************
	// Type conversion
	// ***************
	convertInt := 1
	convertFloat := float32(convertInt)
	fmt.Println(convertFloat, reflect.TypeOf(convertFloat))
	convertAnotherFloat := 1.8
	convertInt = int(convertAnotherFloat)
	fmt.Println(convertInt, reflect.TypeOf(convertInt))

	// parse string to boolean
	convertBool, err := strconv.ParseBool("true")
	fmt.Println(err == nil, convertBool, reflect.TypeOf(convertBool)) // output: true true bool
	// convert string to int
	strConvertInt, err := strconv.Atoi("-42")
	fmt.Println(err == nil, strConvertInt, reflect.TypeOf(strConvertInt)) // output: true -42 int

	// ***************
	// Type assertion
	// ***************

	// poodle.Bark() // undefined error
	p, ok := poodle.(Dog)
	fmt.Println(ok, p.Bark()) // output: true Bolt: Woff woff!

	// ***************
	// Value types vs reference types
	// ***************

	// Value types
	valueType := "value"
	valueTypeCopy := valueType
	valueType = "value changed"
	fmt.Println(valueType)     // output: value
	fmt.Println(valueTypeCopy) // output: value changed

	// Reference types
	referenceType := []string{"dog", "cat"}
	referenceTypeCopy := referenceType
	referenceTypeCopy[1] = "fish"
	fmt.Println(referenceType)     // output: dog fish
	fmt.Println(referenceTypeCopy) // output: dog fish

	// comparing value types
	arr1 := [2]int{1, 2}
	arr2 := [2]int{1, 2}
	fmt.Println(arr1 == arr2) // output: true

	// comparing reference types
	dog1 := Dog{Name: "bolt"}
	dog2 := Dog{Name: "bolt"}
	fmt.Println(dog1 == dog2) // output: false
}

// Composite type: struct
type Owner struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}
type Account struct {
	BSB           int
	AccountNumber int64
	AccountType   string
	Owners        []Owner
}

// Composite type: interface
type Animal interface {
	// function sets
	Eat() string
}

type Dog struct {
	Name string
}

func (d Dog) Eat() string {
	return fmt.Sprintf("%s eats meat", d.Name)
}

func (d Dog) Bark() string {
	return fmt.Sprintf("%s: Woff woff!", d.Name)
}
