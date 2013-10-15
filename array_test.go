package webhelpers

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	var list = []string{"aa", "bb", "cc"}

	if !StringInSlice("bb", list) {
		t.Fatalf(`StringInSlice("bb", ["aa", "bb", "cc"]) want true not false`)
	}

	if StringInSlice("dd", list) {
		t.Fatalf(`StringInSlice("dd", ["aa", "bb", "cc"] want false not true`)
	}
}

func TestStringEqual(t *testing.T) {
	if StringEqual([]string{"a"}, []string{"a", "b"}) {
		t.Fatalf(`StringEqual(["a"], ["a", "b"]) want false not true`)
	}

	if StringEqual([]string{"a", "b"}, []string{"b", "a"}) {
		t.Fatalf(`StringEqual(["a", "b"], ["b", "a"]) want false not true`)
	}

	if !StringEqual([]string{"a", "b"}, []string{"a", "b"}) {
		t.Fatalf(`StringEqual(["a", "b"], ["a", "b"]) want true not false`)
	}
}
