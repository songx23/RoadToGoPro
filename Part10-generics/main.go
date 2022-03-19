package main

import "net/http"

func main() {
	m1 := Magic{"Protection Spell", []string{"Expecto", "Patronum"}}
	m2 := Magic{"Petrol price is crazy", []string{"Expensive", "Petroleum"}}
	err := CustomError{http.StatusNotImplemented, "Panic! Your code is on fire!"}

	GenericsToString(&m1)
	GenericsToString(&m2)
	GenericsToString(&err)
}
