/*
Package stringset defines a special string array. It only allows unique values stored in the array.

Example (what we did in main.go):

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
*/
package stringset
