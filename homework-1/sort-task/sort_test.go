package main

import (
	"testing"
)

const input string = `Hey my love
Goodbye love

Didn't know what time it was the lights were low
I leaned back on my radio
Some cat was layin down some rock 'n' roll, lotta soul, he said
Then the loud sound did seem to fade
Came back like a slow voice on a wave of phase
That weren't no DJ, that was hazy cosmic jive

La, lo, la, la-la, la, la, la
La, la, lo, la-la, la, la, la
Lo, la, la, la-la, la, la, la
La, la, la, la-la, lo, la, la
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, lo
LA, la, la, la-la, la, la, lo

10.5 219
9 1
-1.5`

func TestSortStrings(t *testing.T) {
	correct := `


-1.5
10.5 219
9 1
Came back like a slow voice on a wave of phase
Didn't know what time it was the lights were low
Goodbye love
Hey my love
I leaned back on my radio
LA, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, lo
La, la, la, la-la, lo, la, la
La, la, lo, la-la, la, la, la
La, lo, la, la-la, la, la, la
Lo, la, la, la-la, la, la, la
Some cat was layin down some rock 'n' roll, lotta soul, he said
That weren't no DJ, that was hazy cosmic jive
Then the loud sound did seem to fade`

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: default sort")
	}
}

func TestSortStringsKey3(t *testing.T) {
	correct := `


-1.5
10.5 219
9 1
Goodbye love
I leaned back on my radio
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
La, lo, la, la-la, la, la, la
Lo, la, la, la-la, la, la, la
LA, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, lo
La, la, la, la-la, lo, la, la
Came back like a slow voice on a wave of phase
La, la, lo, la-la, la, la, la
Then the loud sound did seem to fade
Hey my love
That weren't no DJ, that was hazy cosmic jive
Some cat was layin down some rock 'n' roll, lotta soul, he said
Didn't know what time it was the lights were low`

	options := Options{
		keyFlag:         3,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: key -3 sort")
	}
}

func TestSortStringsIgnoreCase(t *testing.T) {
	correct := `


-1.5
10.5 219
9 1
Came back like a slow voice on a wave of phase
Didn't know what time it was the lights were low
Goodbye love
Hey my love
I leaned back on my radio
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
LA, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, lo
La, la, la, la-la, lo, la, la
La, la, lo, la-la, la, la, la
La, lo, la, la-la, la, la, la
Lo, la, la, la-la, la, la, la
Some cat was layin down some rock 'n' roll, lotta soul, he said
That weren't no DJ, that was hazy cosmic jive
Then the loud sound did seem to fade`

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  true,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: ignore case sort")
	}
}

func TestSortStringsUnique(t *testing.T) {
	correct := `
-1.5
10.5 219
9 1
Came back like a slow voice on a wave of phase
Didn't know what time it was the lights were low
Goodbye love
Hey my love
I leaned back on my radio
LA, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, lo
La, la, la, la-la, lo, la, la
La, la, lo, la-la, la, la, la
La, lo, la, la-la, la, la, la
Lo, la, la, la-la, la, la, la
Some cat was layin down some rock 'n' roll, lotta soul, he said
That weren't no DJ, that was hazy cosmic jive
Then the loud sound did seem to fade`

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      true,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: unique sort")
	}
}

func TestSortStringsReverse(t *testing.T) {
	correct := `Then the loud sound did seem to fade
That weren't no DJ, that was hazy cosmic jive
Some cat was layin down some rock 'n' roll, lotta soul, he said
Lo, la, la, la-la, la, la, la
La, lo, la, la-la, la, la, la
La, la, lo, la-la, la, la, la
La, la, la, la-la, lo, la, la
La, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
LA, la, la, la-la, la, la, lo
I leaned back on my radio
Hey my love
Goodbye love
Didn't know what time it was the lights were low
Came back like a slow voice on a wave of phase
9 1
10.5 219
-1.5


`
	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     true,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: reverse sort")
	}
}

func TestSortStringsNumeric(t *testing.T) {
	correct := `-1.5



Came back like a slow voice on a wave of phase
Didn't know what time it was the lights were low
Goodbye love
Hey my love
I leaned back on my radio
LA, la, la, la-la, la, la, lo
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, la
La, la, la, la-la, la, la, lo
La, la, la, la-la, lo, la, la
La, la, lo, la-la, la, la, la
La, lo, la, la-la, la, la, la
Lo, la, la, la-la, la, la, la
Some cat was layin down some rock 'n' roll, lotta soul, he said
That weren't no DJ, that was hazy cosmic jive
Then the loud sound did seem to fade
9 1
10.5 219`

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: true,
	}

	result := SortStrings(input, options)

	if result != correct {
		t.Errorf("Incorrect result: numeric sort")
	}
}
