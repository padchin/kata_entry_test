package main

import (
	"math"
)

func romanToInt(s string) int {
	b := []byte(s)
	if len(b) == 0 {
		return 0
	}
	i := 0
	r := 0
	c := 0
	for i < len(b) && b[i] == 'M' {
		c++
		i++
	}
	r += c * 1000
	c = 0
	if i < len(b)-1 && b[i] == 'C' && b[i+1] == 'M' {
		r += 900
		i += 2
	}
	if i < len(b) && b[i] == 'D' {
		r += 500
		i++
	}
	if i < len(b)-1 && b[i] == 'C' && b[i+1] == 'D' {
		r += 400
		i += 2
	}
	for i < len(b) && b[i] == 'C' {
		c++
		i++
	}
	r += c * 100
	c = 0
	if i < len(b)-1 && b[i] == 'X' && b[i+1] == 'C' {
		r += 90
		i += 2
	}
	if i < len(b) && b[i] == 'L' {
		r += 50
		i++
	}
	if i < len(b)-1 && b[i] == 'X' && b[i+1] == 'L' {
		r += 40
		i += 2
	}
	for i < len(b) && b[i] == 'X' {
		c++
		i++
	}
	r += c * 10
	c = 0
	if i < len(b)-1 && b[i] == 'I' && b[i+1] == 'X' {
		r += 9
		i += 2
	}
	if i < len(b) && b[i] == 'V' {
		r += 5
		i++
	}
	if i < len(b)-1 && b[i] == 'I' && b[i+1] == 'V' {
		r += 4
		i += 2
	}
	for i < len(b) && b[i] == 'I' {
		c++
		i++
	}
	r += c
	return r
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	s := ""
	L := map[int][2]string{
		1: {"I", "V"},
		2: {"X", "L"},
		3: {"C", "D"},
		4: {"M", ""},
	}
	n := int(math.Log10(float64(num))) + 1 // number of digits
	buffer := make([]int, n)
	for t := n - 1; t >= 0; t-- {
		d := int(math.Pow10(t))
		z := num / d
		buffer[t] = z
		num -= z * d
	}
	for i := n; i > 0; i-- {
		d := buffer[i-1]
		switch {
		case d == 9:
			s += L[i][0] + L[i+1][0]
		case d == 4:
			s += L[i][0] + L[i][1]
		case d >= 5:
			s += L[i][1]
			for j := 0; j < d-5; j++ {
				s += L[i][0]
			}
		default:
			for j := 0; j < d; j++ {
				s += L[i][0]
			}
		}
	}
	return s
}
