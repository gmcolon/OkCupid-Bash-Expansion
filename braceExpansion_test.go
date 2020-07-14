package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Input string contains illegal characters
var illegalCharacterInput = []struct {
	name string
	statement string
}{
	{"Numbers and letters", "{A,1,B,2,C,3}"},
	{"Symbols", "{!,@,#,<,?,-,+}"},
	{"Periods, colons, semicolons", "{.,:,;}"},
}

func TestValidateInputMatchesAlphabet(t *testing.T) {
	for _, input := range illegalCharacterInput {
		if os.Getenv("validateInputMatchesAlphabet") == "1" {
			validateInputMatchesAlphabet(input.statement)
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestValidateInputMatchesAlphabet")
		cmd.Env = append(os.Environ(), "validateInputMatchesAlphabet=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			// Input is invalid-- test passes
			continue
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}


// Input string contains empty braces
var emptyBraceInput = []struct {
	name string
	statement string
}{
	{"Single empty", "{}"},
	{"One letter, empty", "{A,}"},
	{"Empty, one letter", "{,A}"},
	{"Multiple letter, empty", "{A,B,C{}}"},
}

func TestValidateInputDoesNotContainEmpty(t *testing.T) {
	for _, input := range emptyBraceInput {
		if os.Getenv("validateInputDoesNotContainEmpty") == "1" {
			validateInputDoesNotContainEmpty(input.statement)
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestValidateInputDoesNotContainEmpty")
		cmd.Env = append(os.Environ(), "validateInputDoesNotContainEmpty=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			// Input is invalid -- test passes
			continue
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}


// Input string contains unbalanced braces
var braceTestInput = []struct {
	name string
	statement string
}{
	{"One close brace at beginning", "}ABC"},
	{"One open brace at beginning", "{ABC"},
	{"Close brace then open brace", "}{"},
	{"Two open braces, one close brace", "{A,B,C{D}"},
	{"One open brace midway through input", "A,{B,C"},
	{"One close brace midway through input", "A,B},C"},
}

func TestCheckValidBraces(t *testing.T) {
	for _, input := range braceTestInput {
		if os.Getenv("validBraces") == "1" {
			validBraces(input.statement)
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestCheckValidBraces")
		cmd.Env = append(os.Environ(), "validBraces=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			// Input is invalid -- test passes
			continue
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)

	}

}

// Full expansion test
var braceExpansionInput = []struct {
	name string
	statement string
	expected string
}{
	{"Three letters in brace", "{A,B,C}", "A B C"},
	{"Side by side braces with letters within", "{A,B}{C,D}", "AC AD BC BD"},
	{"Nested brace with letters", "{A,B{C,D}}","A BC BD"},
	{"One word within braces", "{ABC}", "ABC"},
	{"One word no braces", "ABC", "ABC"},
}

func TestBraceExpansion(t *testing.T) {
	for _, input := range braceExpansionInput {
		actual := start(input.statement)
		if strings.TrimSuffix(input.expected, "\n") != strings.TrimSuffix(actual, "\n") {
			t.Errorf("Wanted: %v, Got: %v", input.expected, actual)
		}
	}
}



