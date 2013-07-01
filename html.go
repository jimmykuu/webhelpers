package webhelpers

import (
	"regexp"
	"strings"
)

var (
	lowerRe, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	brRe, _    = regexp.Compile("<br.*?>")
	tagRe, _   = regexp.Compile("<.*?>")
)

func StripTags(html string) string {
	html = lowerRe.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "&nbsp;", " ", -1)
	html = brRe.ReplaceAllString(html, "")
	html = tagRe.ReplaceAllString(html, "")
	return html
}

// Simplify HTML text by removing tags
func RemoveFormatting(html string) string {
	return StripTags(html)
}
