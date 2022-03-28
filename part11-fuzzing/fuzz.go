package main

func faultyEqual(param1, param2 []byte) bool {
	for i := range param1 {
		if param1[i] != param2[i] {
			return false
		}
	}
	return true
}

func workingEqual(param1, param2 []byte) bool {
	if len(param1) != len(param2) {
		return false
	}

	for i := range param1 {
		if param1[i] != param2[i] {
			return false
		}
	}
	return true
}
