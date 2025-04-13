package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " BYE World",
			expected: []string{"bye", "world"},
		},
		{
			input:    "e C  h   Po    chMa   k  ",
			expected: []string{"e", "c", "h", "po", "chma", "k"},
		},
		// add more cases here
	}

	for k, c := range cases {
		success := true
		actual := cleanInput(c.input)
		lenExp := len(c.expected)
		lenAct := len(actual)
		if lenExp != lenAct {
			success = false
			t.Errorf("Lenght of the case %v (%v) is different from expected %v\n", k, lenAct, lenExp)
		} 

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				success = false
				t.Errorf("The word #%v (%v) in test %v is different from expected %v\n", i, word, k, expectedWord)
			}
		}
		
		if success == true {
			t.Logf("test %v succeeded!\n", k)
		}
	}
}