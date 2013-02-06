package webhelpers

// Truncate text with replacement characters.
func Truncate(text string, length int, indicator string) string {
	r := []rune(text)
	if text == "" {
		return ""
	} else if len(r) <= length {
		return text
	}

	return string(r[:length]) + indicator
}
