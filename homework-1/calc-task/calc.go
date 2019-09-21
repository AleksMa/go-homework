package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Classic top-down parsing with accumulator
// Each function for each expression in Grammar (except Number, Digit, DigitTail)

/*
 * Grammar (BNF)
 *
 * Expr         ::=   Term InnerExpr .
 * InnerExpr    ::= + Term InnerExpr  | - Term InnerExpr | .
 * Term         ::=   Fact InnerTerm .
 * InnerTerm    ::= * Fact InnerTerm  | / Fact InnerTerm  | .
 * Fact			::=   Number | ( Expr ) | - Fact .
 * Number       ::=   Digit DigitTail .
 * Digit        ::=   0 | 1  | ... | 9 .
 * DigitTail    ::=   Digit  | .
 *
 */

const (
	nothing    = iota // End of expression
	positiveOp        // + or *
	negativeOp        // - or /
)

func parser(lexemes []string) (int, error) {
	if len(lexemes) == 0 {
		return 0, errors.New("Empty expression ")
	}
	lexemPointer := 0
	result, err, _ := expr(lexemes, lexemPointer)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func expr(lexemes []string, lexemPointer int) (int, error, int) {
	res, err, leftInc := term(lexemes, lexemPointer)
	if err != nil {
		return 0, err, 0
	}
	next, err, rightInc, status := innerExpr(lexemes, lexemPointer+leftInc)
	if err != nil {
		return 0, err, 0
	}

	if status == positiveOp {
		return res + next, nil, leftInc + rightInc + 1
	} else if status == negativeOp {
		return res + next, nil, leftInc + rightInc + 1
	} else {
		return res, nil, leftInc
	}
}

func innerExpr(lexemes []string, lexemPointer int) (int, error, int, int) {
	if lexemPointer < len(lexemes) && lexemes[lexemPointer] != "+" && lexemes[lexemPointer] != "-" {
		return 0, nil, 0, nothing
		// return 0, fmt.Errorf("Unsupported construction %v [add] ", lexemPointer), 0, nothing
	}
	if lexemPointer >= len(lexemes) {
		return 0, nil, 0, nothing
	}
	status := positiveOp
	if lexemes[lexemPointer] == "-" {
		status = negativeOp
	}

	res, err, leftInc := term(lexemes, lexemPointer+1)
	if err != nil {
		return 0, err, 0, nothing
	}
	if lexemPointer+leftInc >= len(lexemes) || (lexemes[lexemPointer+leftInc] != "+" && lexemes[lexemPointer+leftInc] != "-") {
		return res, nil, leftInc, status
	}
	next, err, rightInc, status := innerExpr(lexemes, lexemPointer+leftInc)
	if err != nil {
		return 0, err, 0, nothing
	}
	if status == positiveOp {
		return res + next, nil, leftInc + rightInc + 1, status
	} else if status == negativeOp {
		return res + next, nil, leftInc + rightInc + 1, status
	} else {
		return res, nil, leftInc, status
	}
}

func term(lexemes []string, lexemPointer int) (int, error, int) {
	res, err, leftInc := factor(lexemes, lexemPointer)
	if err != nil {
		return 0, err, 0
	}
	next, err, rightInc, status := innerTerm(lexemes, lexemPointer+leftInc)
	if err != nil {
		return 0, err, 0
	}

	if status == positiveOp {
		return res * next, nil, leftInc + rightInc + 1
	} else if status == negativeOp {
		return res / next, nil, leftInc + rightInc + 1
	} else {
		return res, nil, leftInc
	}
}

func innerTerm(lexemes []string, lexemPointer int) (int, error, int, int) {
	if lexemPointer < len(lexemes) && lexemes[lexemPointer] != "*" && lexemes[lexemPointer] != "/" {
		return 0, nil, 0, nothing
		//return 0, fmt.Errorf("Unsupported construction %v [mul] ", lexemPointer), 0, nothing
	}
	if lexemPointer >= len(lexemes) {
		return 0, nil, 0, nothing
	}
	status := positiveOp
	if lexemes[lexemPointer] == "/" {
		status = negativeOp
	}

	res, err, leftInc := factor(lexemes, lexemPointer+1)
	if err != nil {
		return 0, err, 0, nothing
	}
	if lexemPointer+leftInc >= len(lexemes) || (lexemes[lexemPointer+leftInc] != "*" && lexemes[lexemPointer+leftInc] != "/") {
		return res, nil, leftInc, status
	}
	next, err, rightInc, status := innerTerm(lexemes, lexemPointer+leftInc)
	if err != nil {
		return 0, err, 0, nothing
	}
	if status == positiveOp {
		return res * next, nil, leftInc + rightInc + 1, status
	} else if status == negativeOp {
		return res / next, nil, leftInc + rightInc + 1, status
	} else {
		return res, nil, leftInc, status
	}
}

func factor(lexemes []string, lexemPointer int) (int, error, int) {
	if num, err := strconv.Atoi(lexemes[lexemPointer]); err == nil {
		return num, nil, 1
	}
	if lexemes[lexemPointer] == "(" {
		res, err, inc := expr(lexemes, lexemPointer+1)
		if err != nil {
			return 0, err, 0
		}
		if lexemes[lexemPointer+1+inc] != ")" {
			return 0, fmt.Errorf("Unexpected symbol. Expected: ')', got: %s ", lexemes[lexemPointer+1+inc]), 0
		}
		return res, nil, inc + 2
	}
	if lexemes[lexemPointer] == "-" {
		res, err, inc := factor(lexemes, lexemPointer)
		if err != nil {
			return 0, err, 0
		}
		return res, nil, inc + 1
	}
	return 0, errors.New("Unexpected construction "), 0
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

	result, err := parser(lexemes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
