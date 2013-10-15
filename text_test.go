package webhelpers

import (
	"testing"
)

type TextAndWant struct {
	Text string
	Want string
}

type TextAndWho struct {
	Text string
	Who  []string
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

var atWhoCases = []TextAndWho{
	TextAndWho{"", []string{}},
	TextAndWho{"@jimmykuu", []string{"jimmykuu"}},
	TextAndWho{" @jimmykuu", []string{"jimmykuu"}},
	TextAndWho{"Hi, @jimmykuu", []string{"jimmykuu"}},
	TextAndWho{"Hi,@jimmykuu", []string{}},
	TextAndWho{"Hi, @jimmykuu, @tom", []string{"tom"}},
	TextAndWho{"Hi, @jimmykuu and @tom and @jimmykuu again", []string{"jimmykuu", "tom"}},
	TextAndWho{"@jimmykuu\nanother line @john", []string{"jimmykuu", "john"}},
	TextAndWho{"jimmykuu@gmail.com", []string{}},
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

func TestAtWho(t *testing.T) {
	for _, textAndWho := range atWhoCases {
		actual := AtWho(textAndWho.Text)
		if !StringEqual(actual, textAndWho.Who) {
			t.Fatalf("AtWho(`%s`) want `%s` not `%s`", textAndWho.Text, textAndWho.Who, actual)
		}
	}
}
