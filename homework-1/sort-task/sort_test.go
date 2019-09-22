package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestSortStrings(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/sorted.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: default sort")
	}
}

func TestSortStringsKey3(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/key3.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         3,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: key -3 sort")
	}
}

func TestSortStringsIgnoreCase(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/ignore_case.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  true,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: ignore case sort")
	}
}

func TestSortStringsUnique(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/unique.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      true,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: unique sort")
	}
}

func TestSortStringsReverse(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/reverse.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     true,
		outputFlag:      "",
		numericSortFlag: false,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: reverse sort")
	}
}

func TestSortStringsNumeric(t *testing.T) {
	data, err := ioutil.ReadFile("tests/starman.txt")
	if err != nil {
		log.Fatal(err)
	}

	correct, err := ioutil.ReadFile("tests/numeric.txt")
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         1,
		ignoreCaseFlag:  false,
		uniqueFlag:      false,
		reverseFlag:     false,
		outputFlag:      "",
		numericSortFlag: true,
	}

	result := SortStrings(string(data), options)

	if result != string(correct) {
		t.Errorf("Incorrect result: numeric sort")
	}
}
