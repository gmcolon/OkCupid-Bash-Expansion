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
	{"One letter in braces", "{A}", "A"},
	{"One word no braces", "ABC", "ABC"},
	{"Three letters in brace", "{A,B,C}", "A B C"},
	{"Side by side braces with letters within", "{A,B}{C,D}", "AC AD BC BD"},
	{"Nested brace with letters", "{A,B{C,D}}","A BC BD"},
	{"One word within braces", "{ABC}", "ABC"},
	{"Doubly nested brace with letters", "{A,B{C,D{E,F}}}","A BC BDE BDF"},
	{"Doubly nested brace with letters", "{A,B{C,D{E{X,Y},F}}}","A BC BDEX BDEY BDF"},
	{"Nested side by side brace", "{A,B{C,D}{E,F}}","A BCE BCF BDE BDF"},
	{"Nested with character before and after inner brace", "{A,B{E,F,G}C,F}", "A BEC BFC BGC F"},
}

func TestRecursiveDescentParser(t *testing.T) {
	for _, input := range expandInput {
		output := Expand(strings.TrimSuffix(input.statement, "\n"))
		if strings.TrimSuffix(input.expected, "\n") != output {
			t.Errorf("Wanted: %v, Got: %v", input.expected, output)
		}
	}
}
