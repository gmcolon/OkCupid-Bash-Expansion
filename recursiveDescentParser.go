package main

import (
	"strings"
	"unicode"
)

// Using a recursive descent parser to split the input & recursively expand it.
// The grammar for the parser is as follows:
//	1. statement = expression, expression  (Comma separated expressions)
//	2. expression = expression_part expression_part (One or more expression parts)
//	3. expression_part = letter | {statement} (A letter OR a statement)


//Initial state
var index = 0
var inputString = ""

// Initial entry into the recursiveDescentParser.
// All inout being passed in should be valid.
func Expand(input string) (resultString string) {
	// Resetting before each run
	index = 0
	inputString = input
	var resultArr = parseStatement()
	resultString = strings.Join(resultArr[:], " ")
	return
}

// A Statement is made up of comma separated expressions
// To parse a statement, we parse each expression and concatenate the results
func parseStatement() (result []string) {
	for index < len(inputString){
		var items = parseExpression()
		result = append(result, items...)
		if index < len(inputString) && inputString[index] == ',' {
			index++
		} else {
			break
		}
	}
	return
}

// An expressions is made up of one or more expression parts
// To parse expressions we parse each expression part and combine them
func parseExpression() (leftItems []string) {
	for index < len(inputString) {
		var rightItems = parseExpressionPart()
		if rightItems == nil {
			break
		}
		if leftItems == nil {
			leftItems = rightItems
		} else {
			leftItems = combine(leftItems, rightItems);
		}
	}
	return
}

// Helper function to combine expression parts
func combine(leftItems []string, rightItems []string) (result []string) {
	for _, leftItem := range leftItems {
		for _, rightItem := range rightItems {
			result = append(result, leftItem + rightItem)
		}
	}
	return
}

// An expression part is made up of letters or a statement
// To parse an expression part,
//  if it's an expression we parse the expression per above
//  if it's a letter, we read the letter and any letters directly after
func parseExpressionPart() (items []string) {
	if index < len(inputString){
		var nextChar = inputString[index]
		if nextChar == '{'{
			//Skip opening brace
			index = index+1
			items = parseStatement()
			//Skip closing brace
			index = index+1
			return
		} else if unicode.IsLetter(rune(nextChar)) {
			var letters = readLetters()
			items = append(items, letters)
		} else {
			return nil
		}
	}

	return
}

// Helper function to read contiguous letters
func readLetters() (letters string) {
	var sb strings.Builder
	if index >= len(inputString){
		return
	}
	for index < len(inputString){
		var nextChar = inputString[index]
		index++
		if unicode.IsLetter(rune(nextChar)){
			sb.WriteRune(rune(nextChar))
		} else {
			index--
			break
		}
	}
	letters = sb.String()
	return
}


