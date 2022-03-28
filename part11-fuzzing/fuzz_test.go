package main

import (
	"testing"
)

func FuzzFaultyEqual(f *testing.F) {
	f.Add([]byte("abc"), []byte("abc"))
	f.Add([]byte("abc"), []byte("123"))

	f.Fuzz(func(t *testing.T, param1, param2 []byte) {
		e := faultyEqual(param1, param2)
		if e != (string(param1) == string(param2)) {
			t.Errorf("equal failed. result: %t, param1: %s, param2: %s", e, param1, param2)
		}
	})
}

func FuzzWorkingEqual(f *testing.F) {
	f.Add([]byte("abc"), []byte("abc"))
	f.Add([]byte("abc"), []byte("123"))

	f.Fuzz(func(t *testing.T, param1, param2 []byte) {
		e := workingEqual(param1, param2)
		if e != (string(param1) == string(param2)) {
			t.Errorf("equal failed. result: %t, param1: %s, param2: %s", e, param1, param2)
		}
	})
}
