package main

import "regexp"

// ansiEscapeRe matches ANSI/VT terminal escape sequences:
//   - CSI sequences: ESC [ ... final-byte  (cursor, color, erase commands)
//   - OSC sequences: ESC ] ... BEL or ST   (title, clipboard injection via OSC 52)
//   - Other two-byte sequences: ESC + single char
var ansiEscapeRe = regexp.MustCompile(
	`\x1b` +
		`(?:` +
		`\[[0-?]*[ -/]*[@-~]` + // CSI: ESC [ param-bytes intermediate-bytes final-byte
		`|\][^\x07\x1b]*(?:\x07|\x1b\\)` + // OSC: ESC ] ... BEL|ST
		`|[^[\]]` + // Other: ESC + one char (not [ or ])
		`)`,
)

// stripANSI removes ANSI/VT terminal escape sequences from s.
func stripANSI(s string) string {
	return ansiEscapeRe.ReplaceAllString(s, "")
}
