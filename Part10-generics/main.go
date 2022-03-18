package main

import "net/http"

func main() {
	m := Magic{"Flying spell", []string{"Wingardium", "Leviosa"}}

	err := CustomError{http.StatusNotImplemented, "Panic! Your code is on fire!"}

	GenericsToString(&m)
	GenericsToString(&err)
}
