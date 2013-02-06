package webhelpers

import (
	"testing"
)

type TextAndWant struct {
	Text string
	Want string
}

var cases = []TextAndWant{
	TextAndWant{"", ""},
	TextAndWant{"abc", "abc"},
	TextAndWant{"abcd", "abcd"},
	TextAndWant{"abcdef", "abcd..."},
	TextAndWant{"中文", "中文"},
	TextAndWant{"这是一段中文", "这是一段..."},
	TextAndWant{"Hello, 世界", "Hell..."},
	TextAndWant{"你好，World", "你好，W..."},
}

func TestTruncate(t *testing.T) {
	length := 4
	for _, textAndWant := range cases {
		actual := Truncate(textAndWant.Text, length, "...")
		if actual != textAndWant.Want {
			t.Fatalf("Truncate(`%s`) want `%s` not `%s`", textAndWant.Text, textAndWant.Want, actual)
		}
	}
}
