package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// Classic top-down parsing with accumulator
// Each function for each expression in Grammar (except Number)

/*
 * Grammar (BNF)
 *
 * Expr         ::=   Term InnerExpr .
 * InnerExpr    ::= + Term InnerExpr  | - Term InnerExpr .
 * Term         ::=   Fact InnerTerm .
 * InnerTerm    ::= * Fact InnerTerm  | / Fact InnerTerm  | .
 * Fact			::=   Number | ( Expr ) | - Fact .
 * Number       ::=   Digit DigitTail .
 * Digit        ::=   0 | 1  | ... | 9 .
 * DigitTail    ::=   Digit  | .
 *
 */

func lexer(expression string) ([]string, error) {
	lexemes := make([]string, 0)
	acc := ""

	for _, ch := range expression {
		if ch == ' ' {
			if acc != "" {
				lexemes = append(lexemes, acc)
				acc = ""
			}
			continue
		} else if unicode.IsDigit(ch) {
			acc = acc + string(ch)
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '(' || ch == ')' {
			if acc != "" {
				lexemes = append(lexemes, acc)
				acc = ""
			}
			lexemes = append(lexemes, string(ch))
		} else {
			return nil, errors.New("Unexpected symbol ")
		}
	}
	if acc != "" {
		lexemes = append(lexemes, acc)
	}
	return lexemes, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: \"go run calc.go expression\"")
		return
	}

	/*	expressionParts := os.Args[1:]
		var expressionBuilder strings.Builder
		for _, part := range expressionParts {
			expressionBuilder.WriteString(part)
		}
		fmt.Println(expressionBuilder.String())*/

	expression := os.Args[1]
	lexemes, err := lexer(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", lexemes)

}
