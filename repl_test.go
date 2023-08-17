package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "PHIL SAMOEI",
			expected: []string{
				"phil",
				"samoei",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", actual, cs.expected)
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v is not equal %v", actualWord, expectedWord)
			}
		}

	}

}
