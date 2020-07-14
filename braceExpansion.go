package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const emptyValues = "{[a-zA-z]*,*}"


// This program implements a similar behavior to bash's brace expansion.
// For a valid input, it will expand the expression and return it as a string.
// For invalid input, it will print nothing and exit.

// Input characters are restricted to [a-zA-Z{},]
// An example invocation of the program will be:
// 	$ echo "{A,B,C}" | okCupidTakeHome
// And it should print:
// 	A B C

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		os.Exit(1)
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		// Was not invoked using a pipe
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	start(strings.TrimSpace(text))
}

func start(text string) (expanded string){
	// validate input
	validateInputMatchesAlphabet(text)
	validateInputDoesNotContainEmpty(text)
	validBraces(text)

	//expand input
	expanded = strings.TrimSpace(Expand(text))
	fmt.Println(expanded)
	return
}

// Below we have helper methods to validate input

// Check that the input matches the allowed characters:
// ASCII alpha characters, braces, and comma.
func validateInputMatchesAlphabet(toValidate string) {
	for _, r := range toValidate {
		isNotLetter := (r < 'a' || r > 'z') && (r < 'A' || r > 'Z')
		isNotBraces := (r != '{') && (r != '}')
		isNotComma := r != ','
		if isNotLetter && isNotBraces && isNotComma {
			// If character is not any, fail
			os.Exit(1)
		}
	}
}

// Check that the input has no empty values:
// Examples: {}, {A,}, {,}
func validateInputDoesNotContainEmpty(toValidate string) {
	for i, r := range toValidate {
		if r == '{' {
			if i < len(toValidate){
				cur := toValidate[i+1]
				if  cur == ',' || cur == '}' {
					os.Exit(1)
				}
			}
		} else if r == '}'{
			if i > 0 {
				cur := toValidate[i-1]
				if  cur == ',' {
					os.Exit(1)
				}
			}
		}
	}
}

// Check that braces are balanced
func validBraces(toValidate string){
	var stack []string
	for _, c := range toValidate {
		charVal := string(c)
		if charVal == "{" {
			//Push opening brace
			stack = append(stack, "{")
		} else if charVal == "}" {
			if len(stack) > 0 {
				// Pop corresponding brace
				n := len(stack) - 1
				stack = stack[:n]
			} else {
				// Mismatched braces
				os.Exit(1)
			}
		}
	}
	// Braces were not matched -- exit program
	if len(stack) > 0{
		os.Exit(1)
	}
}