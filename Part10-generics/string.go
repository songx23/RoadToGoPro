package main

func StringSliceContains(src []string, trg string) bool {
	for _, s := range src {
		if s == trg {
			return true
		}
	}
	return false
}
