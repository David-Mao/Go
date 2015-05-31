package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct{ input, expect_output string }{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		output := Reverse(c.input)
		if output != c.expect_output {
			t.Errorf("Reverse(%q) should be %q, not %q",
				c.input, c.expect_output, output)
		}
	}
}
