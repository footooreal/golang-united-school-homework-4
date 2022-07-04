package string_sum

import (
	"errors"
	"fmt"
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

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")
	operArr := []int{}
	symArr := []byte{}

	for i, v := range input {
		x, err := strconv.Atoi(string(v))
		if input[i-1] != 0 {
			symArr = append(symArr, input[i-1])
		} else {
			symArr = append(symArr, 0)
		}
		if err == nil {
			operArr = append(operArr, x)
		}
	}

	if len(operArr) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	sym1 := string(symArr[0])
	sym2 := string(symArr[1])
	oper1 := operArr[0]
	oper2 := operArr[1]
	if sym1 == "-" {
		oper1 = -oper1
	}
	if sym2 == "-" {
		oper2 = -oper2
	}

	answer := oper1 + oper2

	return strconv.Itoa(answer), nil

}
