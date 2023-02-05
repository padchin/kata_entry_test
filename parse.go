package main

import (
	"regexp"
	"strconv"
	"strings"
)

func validateQuery(query string) (operand1 int, operand2 int, operator string, roman bool, err error) {
	// remove spaces
	query = strings.Join(strings.Fields(query), "")
	query = strings.ToUpper(query)
	//search for operator
	seps := []string{"+", "-", "/", "*"}
	found := false
	var sOperand1 string
	var sOperand2 string
	for _, sep := range seps {
		if strings.Contains(query, sep) {
			ops := strings.Split(query, sep)
			if len(ops) == 2 {
				sOperand1 = ops[0]
				sOperand2 = ops[1]
				operator = sep
				found = true
				break
			} else {
				return 0, 0, "", false, ErrInvalidOperator
			}
		}
	}
	if !found {
		return 0, 0, "", false, ErrInvalidOperator
	}
	bOperand1Roman := isRoman(sOperand1)
	bOperand1Numeric := isNumeric(sOperand1)
	bOperand2Roman := isRoman(sOperand2)
	bOperand2Numeric := isNumeric(sOperand2)
	if bOperand1Numeric && bOperand2Numeric {
		operand1, _ = strconv.Atoi(sOperand1)
		operand2, _ = strconv.Atoi(sOperand2)
		return operand1, operand2, operator, false, nil
	}
	if bOperand1Roman && bOperand2Roman {
		operand1 = romanToInt(sOperand1)
		operand2 = romanToInt(sOperand2)
		return operand1, operand2, operator, true, nil
	}
	if (bOperand1Numeric && bOperand2Roman) || (bOperand2Numeric && bOperand1Roman) {
		return 0, 0, "", false, ErrMixedNumerals
	}
	return 0, 0, "", false, ErrCommonFormat
}

func isNumeric(query string) bool {
	if query == "" {
		return false
	}
	check, _ := regexp.Match("^[0-9]+$", []byte(query))
	return check
}

func isRoman(query string) bool {
	if query == "" {
		return false
	}
	check, _ := regexp.Match("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", []byte(query))
	return check
}
