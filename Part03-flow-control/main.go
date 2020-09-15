package main

import (
	"fmt"
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
		fmt.Println(newPlace, " is a valid location on Earth.")
	}
	// errors as the variables are out of scope
	// fmt.Println(anotherPlace) // undefined: anotherPlace
	// fmt.Println(newPlace)     // undefined: newPlace

	// switch statement
	numText := "one"
	switch numText {
	case "one":
		fmt.Println(numText, " = ", 1)
	case "two":
		fmt.Println(numText, " = ", 2)
	default:
		fmt.Println(numText, " is not 1 or 2")
	}

	// generic switch statement
	switch {
	case "one" == "1":
		fmt.Println("This is impossible")
	case dailyNewCOVIDCases < covidCaseLimit && time.Now().After(time.Date(2020, 9, 28, 0, 0, 0, 0, loc)):
		fmt.Println("Easing lockdown")
	default:
		fmt.Println("Catching all other scenarios")
	}
}
