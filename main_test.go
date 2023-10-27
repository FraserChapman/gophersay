package main

import (
	"testing"
)

func TestMaxLen(t *testing.T) {
	cases := []struct {
		input    []string
		expected int
	}{
		{[]string{""}, 0},
		{[]string{"foo", "bar", "baz"}, 3},
		{[]string{"hello", "world"}, 5},
		{[]string{"a", "bb", "ccc", "dddd"}, 4},
	}

	for _, c := range cases {
		result := maxLen(c.input...)
		if result != c.expected {
			t.Errorf("maxLen(%v) == %d, expected %d", c.input, result, c.expected)
		}
	}
}

func TestGopherSays(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			[]string{},
			"+--------------+\n| Don't panic! |\n+--------------+\n" + gopher,
		},
		{
			[]string{"Hello, World!", "Love gopher x"},
			"+---------------+\n| Hello, World! |\n| Love gopher x |\n+---------------+\n" + gopher,
		},
	}

	for _, c := range cases {
		result := gopherSay(c.input...)
		if result != c.expected {
			t.Errorf("gopherSays(%v) == \n%s\nExpected:\n%s\n", c.input, result, c.expected)
		}
	}
}
