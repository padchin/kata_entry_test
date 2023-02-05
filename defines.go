package main

import "errors"

var ErrOutOfBounds = errors.New("entered values outside the specified range [1..10]")
var ErrMixedNumerals = errors.New("cannot use Roman numerals and Arabic numerals in the same query")
var ErrRomanLessThanOne = errors.New("roman numerals cannot be less than one")
var ErrInvalidOperator = errors.New("invalid operator")
var ErrCommonFormat = errors.New("the query must consist of two integer numbers and one operator for them")
