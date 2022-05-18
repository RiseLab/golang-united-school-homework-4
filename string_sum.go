package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

const regexTerms = `[+-]?[^+-]*`

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if input == "" {
		return "", fmt.Errorf("parsing error: %w", errorEmptyInput)
	}

	reTerms := regexp.MustCompile(regexTerms)
	terms := reTerms.FindAllString(input, -1)

	if len(terms) != 2 {
		return "", fmt.Errorf("parsing error: %w", errorNotTwoOperands)
	}

	sum := 0
	for _, term := range terms {
		v, e := strconv.Atoi(term)
		if e != nil {
			return "", fmt.Errorf("parsing error: %w", e)
		}
		sum += v
	}

	return strconv.Itoa(sum), nil
}
