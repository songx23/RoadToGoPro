package main

func IntSliceContains(src []int, trg int) bool {
	for _, s := range src {
		if s == trg {
			return true
		}
	}
	return false
}
