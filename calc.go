package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		query = query[:len(query)-2]
		a, b, operand, roman, err := validateQuery(query)
		if err != nil {
			log.Fatal(err)
		}
		if a > 10 || a < 1 || b > 10 || b < 1 {
			log.Fatal(ErrOutOfBounds)
		}
		if !roman {
			switch operand {
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
			switch operand {
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
