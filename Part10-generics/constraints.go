package main

import (
	"fmt"
	"strings"
)

// Generics constraint using an interface

type Stringer interface {
	String() string
}

type Magic struct {
	Name  string
	Spell []string
}

func (m *Magic) String() string {
	return fmt.Sprintf("This is a %s. Repeat after me: %s", m.Name, strings.Join(m.Spell, " "))
}

type CustomError struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func (e *CustomError) String() string {
	return fmt.Sprintf("Error: status code %d, error message %s", e.HTTPStatusCode, e.ErrorMessage)
}

func GenericsToString[Str Stringer](structure Str) {
	fmt.Println(structure.String())
}
