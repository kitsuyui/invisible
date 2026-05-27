package main

import "testing"

func TestStripANSI(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "plain text unchanged",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "CSI color sequence",
			input: "\x1b[31mred\x1b[0m",
			want:  "red",
		},
		{
			name:  "CSI erase screen",
			input: "before\x1b[2Jafter",
			want:  "beforeafter",
		},
		{
			name:  "CSI cursor hide",
			input: "\x1b[?25l",
			want:  "",
		},
		{
			name:  "OSC terminal title",
			input: "\x1b]0;evil title\x07",
			want:  "",
		},
		{
			name:  "OSC with ST terminator",
			input: "\x1b]0;evil\x1b\\text",
			want:  "text",
		},
		{
			name:  "OSC 52 clipboard injection",
			input: "\x1b]52;c;cGVybA==\x07",
			want:  "",
		},
		{
			name:  "two-byte ESC sequence",
			input: "\x1b=",
			want:  "",
		},
		{
			name:  "mixed content",
			input: "safe\x1b[31mcolor\x1b[0msafe",
			want:  "safecolorsafe",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stripANSI(tc.input)
			if got != tc.want {
				t.Errorf("stripANSI(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
