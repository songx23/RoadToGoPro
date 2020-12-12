package main

import (
	"fmt"

	"github.com/songx23/RoadToGoPro/part05/pkg/stringset"
)

func main() {
	// Creating a new string set
	set := stringset.NewUniqueStringSet("first string", "first string")
	fmt.Println(set)  // output: [first string]
	set.Add("second string")
	fmt.Println(set)  // output: [first string second string]
	set.Add("second string")
	fmt.Println(set)  // output: [first string second string]
	set.Remove("first string")
	fmt.Println(set)  // output: [second string]
	set.Format()
	fmt.Println(set)  // output: [Second String]
}
