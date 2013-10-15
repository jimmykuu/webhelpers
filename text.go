package webhelpers

import (
	"regexp"
	"strings"
)

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

// find the usernames beed ated.
func AtWho(text string) []string {
	var result = []string{}
	var username string
	usernameRe, _ := regexp.Compile(`^[a-zA-Z0-9_]+$`)
	for _, line := range strings.Split(text, "\n") {
		if len(line) == 0 {
			continue
		}

		for {
			index := strings.Index(line, "@")
			if index == -1 {
				break
			} else if index > 0 {
				if line[index-1] == ' ' {
					line = line[index:]
				} else {
					line = line[index+1:]
				}
			} else if index == 0 {
				// the "@" is first characters
				endIndex := strings.Index(line, " ")
				if endIndex == -1 {
					username = line[1:]
				} else {
					username = line[1:endIndex]
				}

				if username != "" && usernameRe.MatchString(username) && !StringInSlice(username, result) {
					result = append(result, username)
				}

				if endIndex == -1 {
					break
				}

				line = line[endIndex:]
			}
		}
	}

	return result
}
