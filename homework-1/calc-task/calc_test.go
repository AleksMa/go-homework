package main

import (
	"fmt"
	"testing"
)

func incorrectResult(result int, correctResult int) string {
	fmt.Println(result)
	return fmt.Sprintf("Incorrect result. Expected: %v. Got: %v\n", correctResult, result)
}

func TestCalculateAdd(t *testing.T) {
	correctResult := 5
	expression := "2+3"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateSub(t *testing.T) {
	correctResult := -1
	expression := "2 - 3"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateMul(t *testing.T) {
	correctResult := 6
	expression := "2* 3"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateDiv(t *testing.T) {
	correctResult := 0
	expression := "2/3"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateUnaryMinus(t *testing.T) {
	correctResult := -10
	expression := "5 * (-2)"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateBrakes(t *testing.T) {
	correctResult := 5
	expression := "5 * (3-2)"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateIncorrectExpression(t *testing.T) {
	expression := "5 * a"
	result, err := Calculate(expression)
	if err == nil {
		t.Errorf("Expected: lexem validation. Got: %v\n", result)
	}
}

func TestCalculateChaining(t *testing.T) {
	correctResult := 22
	expression := "5 * 2 + (10 - (-1) * 2)"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}

func TestCalculateSophisticated(t *testing.T) {
	correctResult := 0
	expression := "(20 - - 14 - (4 - 7))/(13*6 - (100 - - 2*6))*(8 + - 18 + - 4*1 + (-9)*4 - 11*(20 - 20))/(-2/ - 12/(10* - 9)*(-18/ - 17 - (12 - (100 + 4))) + - 18*8 - - 19*(6 + 5) - 13/ - 2* - 7*5* - 10)"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != correctResult {
		t.Errorf(incorrectResult(result, correctResult))
	}
}
