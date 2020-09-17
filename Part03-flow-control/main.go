package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// if...else... statement example
	grocery := []string{"onion", "milk"}
	if len(grocery) < 2 {
		fmt.Println("Are you sure you have included everything you need in the shopping list?")
	} else if len(grocery) < 20 {
		fmt.Println("Looks good. Time to go to the supermarket.")
	} else {
		fmt.Println("You have too many things in your shopping list.")
	}
	// output: Are you sure you have included everything you need in the shopping list.

	// condition with logic gate
	covidCaseLimit := 50
	dailyNewCOVIDCases := 45
	loc, _ := time.LoadLocation("Australia/Melbourne")
	if dailyNewCOVIDCases < covidCaseLimit && time.Now().After(time.Date(2020, 9, 28, 0, 0, 0, 0, loc)) {
		fmt.Println("Starting to lift restrictions! :)")
	} else {
		fmt.Println("Nope, still in lockdown :(")
	}
	// output: depends on when you run the program, go have a try yourself ;)

	if location, err := time.LoadLocation("Mars/Somewhere"); err != nil {
		fmt.Println(err) // unknown time zone Mars/Somewhere
		fmt.Println("This is not a valid location on Earth.", " Default time zone:", location)
		// output: This is not a valid location on Earth.  Default time zone: UTC
	}
	// error if trying to print out location
	// fmt.Println(location) // undefined: location

	// variable scopes
	if _, err := time.LoadLocation("Mars/Somewhere"); err != nil {
		anotherPlace := "Asgard"
		fmt.Println("This is not a valid location on Earth. Choosing somewhere else... ", anotherPlace)
	} else {
		newPlace := "New Asgard"
		fmt.Println(newPlace, "is a valid location on Earth.")
	}
	// errors as the variables are out of scope
	// fmt.Println(anotherPlace) // undefined: anotherPlace
	// fmt.Println(newPlace)     // undefined: newPlace

	// switch statement
	numText := "one"
	switch numText {
	case "one":
		fmt.Println(numText, "=", 1)
	case "two":
		fmt.Println(numText, "=", 2)
	default:
		fmt.Println(numText, "is not 1 or 2")
	}
	//output: one = 1

	// generic switch statement
	switch {
	case "one" == "1":
		fmt.Println("This is impossible")
	case dailyNewCOVIDCases < covidCaseLimit && time.Now().After(time.Date(2020, 9, 28, 0, 0, 0, 0, loc)):
		fmt.Println("Easing lockdown")
	default:
		fmt.Println("Catching all other scenarios")
	}
	// output: depends on when you execute this program

	// for loop: basic form
	lockdownLevels := []string{"stage1", "stage2", "stage3", "stage4"}
	for i := 0; i < len(lockdownLevels); i++ {
		switch {
		case strings.Contains(lockdownLevels[i], "1"):
			fmt.Println("In", lockdownLevels[i], "lockdown. Life is great.")
		case strings.Contains(lockdownLevels[i], "2"):
			fmt.Println("In", lockdownLevels[i], "lockdown. Life is good.")
		case strings.Contains(lockdownLevels[i], "3"):
			fmt.Println("In", lockdownLevels[i], "lockdown. Life is not too bad.")
		case strings.Contains(lockdownLevels[i], "4"):
			fmt.Println("In", lockdownLevels[i], "lockdown. Life is not too good.")
		default:
			fmt.Println("Unknown lockdown stage")
		}
	}

	// for loop: while form
	whileIndex := 0
	for whileIndex < len(lockdownLevels) {
		switch {
		case strings.Contains(lockdownLevels[whileIndex], "1"):
			fmt.Println("In", lockdownLevels[whileIndex], "lockdown. Life is great.")
		case strings.Contains(lockdownLevels[whileIndex], "2"):
			fmt.Println("In", lockdownLevels[whileIndex], "lockdown. Life is good.")
		case strings.Contains(lockdownLevels[whileIndex], "3"):
			fmt.Println("In", lockdownLevels[whileIndex], "lockdown. Life is not too bad.")
		case strings.Contains(lockdownLevels[whileIndex], "4"):
			fmt.Println("In", lockdownLevels[whileIndex], "lockdown. Life is not too good.")
		default:
			fmt.Println("Unknown lockdown stage")
		}
		whileIndex++
	}

	// for loop: endless loop
	go func() {
		for {
			fmt.Println("I'm an endless loop.")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
	// output: 5 rows of I'm an endless loop.

	// for loop: range operator
	for index, value := range lockdownLevels {
		fmt.Print("Index ", index, ": ")
		switch {
		case strings.Contains(value, "1"):
			fmt.Println("In", value, "lockdown. Life is great.")
		case strings.Contains(value, "2"):
			fmt.Println("In", value, "lockdown. Life is good.")
		case strings.Contains(value, "3"):
			fmt.Println("In", value, "lockdown. Life is not too bad.")
		case strings.Contains(value, "4"):
			fmt.Println("In", value, "lockdown. Life is not too good.")
		default:
			fmt.Println("Unknown lockdown stage")
		}
	}

	// for loop: range operator over map
	postcodes := map[int]string{
		3000: "Melbourne",
		2000: "Sydney",
		6000: "Perth",
	}
	for key, value := range postcodes {
		fmt.Println("Postcode:", key, "Name:", value)
	}

	// for loop: range operator over string
	helloString := "Hello!"
	for _, u := range helloString {
		// using fmt package to format the output
		fmt.Printf("Unicode: %U; Character: %c\n", u, u)
	}

	// continue
	for key, value := range postcodes {
		if key == 6000 {
			continue // skip the following code for current iteration
		}
		fmt.Printf("Postcode: %d Name: %s\n", key, value)
	}
	// output:
	//  Postcode: 2000 Name: Sydney
	//  Postcode: 3000 Name: Melbourne

	// break
	for key, value := range postcodes {
		if key == 6000 {
			break // skip the entire loop
		}
		fmt.Printf("Postcode: %d Name: %s\n", key, value)
	}
	// output:
	//  Postcode: 2000 Name: Sydney

	// goto & label
	fmt.Println("First line of code.")
	goto SKIP
	fmt.Println("Second line of code.") // this line is skipped
SKIP:
	fmt.Println("Third line of code.")
	// output:
	//  First line of code.
	//  Third line of code.

	// defer
	dText := "defer #1"
	defer fmt.Println(dText)
	dText = "defer #2"
	defer fmt.Println(dText)
	// do something else
	dText = "not deferred"
	fmt.Println(dText)

	// basic usage of formatter
	whoami := "formatter"
	fmt.Printf("Hello, I'm your helpful %s\n", whoami) // output: Hello, I'm your helpful formatter

	// format verbs: boolean
	booleanTrue := true
	fmt.Printf("Boolean in text: %t\n", booleanTrue) // output: Boolean in text: true

	// format verbs: numbers
	smallInt := 111
	smallFloat := 111.11
	largeFloat := 111111111111111.0
	fmt.Printf(`Binary number: %b
Base 10 number: %d
Unicode: %U
Decimal point: %f
Scientific notationL %e
`, smallInt, smallInt, smallInt, smallFloat, largeFloat)
	// output:
	//  Binary number: 1101111
	//  Base 10 number: 111
	//  Unicode: U+006F
	//  Decimal point: 111.110000
	//  Scientific notationL 1.111111e+14

	// format verbs: string
	simpleString := "I am a happy formatter"
	quotedString := `"I am a happy double quoted formatter"`
	fmt.Printf(`Plain string: %s
Quoted string: %q
Safely escaped quoted string: %q
`, simpleString, simpleString, quotedString)
	// output:
	//  Plain string: I am a happy formatter
	//  Quoted string: "I am a happy formatter"
	//  Safely escaped quoted string: "\"I am a happy double quoted formatter\""

	// format verbs: general
	simpleSlice := []string{"1", "2", "3", "Go"}
	simpleMap := map[string]string{
		"%v":  "value in default format",
		"%+v": "value + field names",
	}
	simpleStruct := Person{
		Name:   "TesterUser",
		Gender: "Secret",
		DoB:    time.Date(1970, 1, 1, 0, 0, 0, 0, loc),
		Age:    50,
		Hobby:  []string{"Go", "Programming", "Break stuff", "Push buttons"},
	}
	fmt.Printf(`Default format of slice: %v
Default format of map: %v
Default format of struct: %v
struct with field names: %+v
Go-syntax representation of slice: %#v
Go-syntax representation of map: %#v
Go-syntax representation of struct: %#v
`, simpleSlice, simpleMap, simpleStruct, simpleStruct, simpleSlice, simpleMap, simpleStruct)
	// output:
	//  Default format of slice: [1 2 3 Go]
	//  Default format of map: map[%+v:value + field names %v:value in default format]
	//  Default format of struct: {TesterUser Secret 1970-01-01 00:00:00 +1000 AEST 50 [Go Programming Break stuff Push buttons]}
	//  struct with field names: {Name:TesterUser Gender:Secret DoB:1970-01-01 00:00:00 +1000 AEST Age:50 Hobby:[Go Programming Break stuff Push buttons]}
	//  Go-syntax representation of slice: []string{"1", "2", "3", "Go"}
	//  Go-syntax representation of map: map[string]string{"%+v":"value + field names", "%v":"value in default format"}
	//  Go-syntax representation of struct: main.Person{Name:"TesterUser", Gender:"Secret", DoB:time.Time{wall:0x0, ext:62135560800, loc:(*time.Location)(0xc0000261e0)}, Age:50, Hobby:[]string{"Go", "Programming", "Break stuff", "Push buttons"}}
}

type Person struct {
	Name   string
	Gender string
	DoB    time.Time
	Age    int
	Hobby  []string
}
