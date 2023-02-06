package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to calculator for Romans and Arabic numbers.")
	for {
		fmt.Print(">")
		query, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading standard input: %v", err)
		}
		query = strings.Replace(query, " ", "", -1)
		query = strings.Replace(query, "\t", "", -1)
		query = strings.Replace(query, "\n", "", -1)
		query = strings.Replace(query, "\r", "", -1)
		query = strings.ToUpper(query)
		a, b, operator, roman, err := validateQuery(query)
		if err != nil {
			log.Fatal(err)
		}
		if a > 10 || a < 1 || b > 10 || b < 1 {
			log.Fatal(ErrOutOfBounds)
		}
		if !roman {
			switch operator {
			case "+":
				fmt.Println(a + b)
			case "-":
				fmt.Println(a - b)
			case "*":
				fmt.Println(a * b)
			case "/":
				fmt.Println(a / b)
			}
		} else {
			switch operator {
			case "+":
				fmt.Println(intToRoman(a + b))
			case "-":
				if a-b <= 0 {
					log.Fatal(ErrRomanLessThanOne)
				}
				fmt.Println(intToRoman(a - b))
			case "*":
				fmt.Println(intToRoman(a * b))
			case "/":
				if a/b == 0 {
					log.Fatal(ErrRomanLessThanOne)
				}
				fmt.Println(intToRoman(a / b))
			}
		}
	}
}
