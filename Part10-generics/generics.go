package main

type ContainsElement interface {
	int | string
}

func SliceContains[C ContainsElement](src []C, trg C) bool {
	for _, s := range src {
		if s == trg {
			return true
		}
	}
	return false
}
