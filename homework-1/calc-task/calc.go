package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Classic top-down parsing
// Each function for each expression in Grammar (except Number, Digit, DigitTail)

/*
 * Grammar (BNF)
 *
 * Expr         ::=   Term InnerExpr .
 * InnerExpr    ::= + Term InnerExpr  | - Term InnerExpr | .
 * Term         ::=   Fact InnerTerm .
 * InnerTerm    ::= * Fact InnerTerm  | / Fact InnerTerm  | .
 * Factor	 	::=   Number | ( Expr ) | - Fact .
 * Number       ::=   Digit DigitTail .
 * Digit        ::=   0 | 1  | ... | 9 .
 * DigitTail    ::=   Digit  | .
 *
 */

func parser(lexemes []string) (int, error) {
	if len(lexemes) == 0 {
		return 0, errors.New("Empty expression ")
	}
	lexemePointer := 0
	result, _, err := expr(lexemes, lexemePointer)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func expr(lexemes []string, lexemePointer int) (int, int, error) {
	res, leftInc, err := term(lexemes, lexemePointer)
	if err != nil {
		return 0, 0, err
	}
	next, rightInc, err := innerExpr(lexemes, lexemePointer+leftInc, res)
	if err != nil {
		return 0, 0, err
	}
	return next, leftInc + rightInc, err
}

func innerExpr(lexemes []string, lexemePointer int, left int) (int, int, error) {
	if lexemePointer < len(lexemes) && lexemes[lexemePointer] != "+" && lexemes[lexemePointer] != "-" {
		return left, 0, nil
	}

	if lexemePointer >= len(lexemes) {
		return left, 0, nil
	}

	res, leftInc, err := term(lexemes, lexemePointer+1)
	if err != nil {
		return 0, 0, err
	}

	if lexemes[lexemePointer] == "+" {
		left = left + res
	} else {
		left = left - res
	}

	next, rightInc, err := innerExpr(lexemes, lexemePointer+leftInc+1, left)
	if err != nil {
		return 0, 0, err
	}
	return next, leftInc + 1 + rightInc, nil
}

func term(lexemes []string, lexemePointer int) (int, int, error) {
	res, leftInc, err := factor(lexemes, lexemePointer)
	if err != nil {
		return 0, 0, err
	}
	next, rightInc, err := innerTerm(lexemes, lexemePointer+leftInc, res)
	if err != nil {
		return 0, 0, err
	}
	return next, leftInc + rightInc, nil
}

func innerTerm(lexemes []string, lexemePointer int, left int) (int, int, error) {
	if lexemePointer < len(lexemes) && lexemes[lexemePointer] != "*" && lexemes[lexemePointer] != "/" {
		return left, 0, nil
	}
	if lexemePointer >= len(lexemes) {
		return left, 0, nil
	}

	res, leftInc, err := factor(lexemes, lexemePointer+1)
	if err != nil {
		return 0, 0, err
	}

	if lexemes[lexemePointer] == "*" {
		left = left * res
	} else {
		left = left / res
	}

	next, rightInc, err := innerTerm(lexemes, lexemePointer+leftInc+1, left)
	if err != nil {
		return 0, 0, err
	}
	return next, leftInc + 1 + rightInc, nil
}

func factor(lexemes []string, lexemePointer int) (int, int, error) {
	if num, err := strconv.Atoi(lexemes[lexemePointer]); err == nil {
		return num, 1, err
	}
	if lexemes[lexemePointer] == "(" {
		res, inc, err := expr(lexemes, lexemePointer+1)
		if err != nil {
			return 0, 0, err
		}
		if lexemes[lexemePointer+1+inc] != ")" {
			return 0, 0, fmt.Errorf("Unexpected symbol. Expected: ')', got: %s ", lexemes[lexemePointer+1+inc])
		}
		return res, inc + 2, nil
	}
	if lexemes[lexemePointer] == "-" {
		res, inc, err := factor(lexemes, lexemePointer+1)
		if err != nil {
			return 0, 0, err
		}
		return -res, inc + 1, nil
	}
	return 0, 0, errors.New("Unexpected construction ")
}

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
			return nil, fmt.Errorf("Unexpected symbol: %v\n ", ch)
		}
	}
	if acc != "" {
		lexemes = append(lexemes, acc)
	}
	return lexemes, nil
}

func Calculate(expression string) (int, error) {
	lexemes, err := lexer(expression)
	if err != nil {
		return 0, err
	}

	result, err := parser(lexemes)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: \"go run calc.go expression\"")
		return
	}

	result, err := Calculate(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
