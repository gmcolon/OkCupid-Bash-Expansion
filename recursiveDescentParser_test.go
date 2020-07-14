package main

import (
	"strings"
	"testing"
)

// Full expansion test
var expandInput = []struct {
	name string
	statement string
	expected string
}{
	{"1. BLAH", "{A,B,C,D{E,F,G}}", "A B C DE DF DG"},
	{"2. One letter in braces", "{A}", "A"},
	{"3. One word no braces", "ABC", "ABC"},
	{"4. Three letters in brace", "{A,B,C}", "A B C"},
	{"5. Side by side braces with letters within", "{A,B}{C,D}", "AC AD BC BD"},
	{"6. Nested brace with letters", "{A,B{C,D}}","A BC BD"},
	{"7. One word within braces", "{ABC}", "ABC"},
	{"8. Doubly nested brace with letters", "{A,B{C,D{E,F}}}","A BC BDE BDF"},
	{"9. Doubly nested brace with letters", "{A,B{C,D{E{X,Y},F}}}","A BC BDEX BDEY BDF"},
	{"10. Nested side by side brace", "{A,B{C,D}{E,F}}","A BCE BCF BDE BDF"},
}

func TestRecursiveDescentParser(t *testing.T) {
	for _, input := range expandInput {
		output := Expand(strings.TrimSuffix(input.statement, "\n"))
		if strings.TrimSuffix(input.expected, "\n") != output {
			t.Errorf("Wanted: %v, Got: %v", input.expected, output)
		}
	}
}
