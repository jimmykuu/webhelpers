package webhelpers

func StringInSlice(str string, list []string) bool {
	for _, element := range list {
		if element == str {
			return true
		}
	}

	return false
}

func StringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}
