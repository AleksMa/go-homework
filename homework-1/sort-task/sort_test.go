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

	result := SortStrings(string(data), 1, false, false, false, false)

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

	result := SortStrings(string(data), 3, false, false, false, false)

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

	result := SortStrings(string(data), 1, true, false, false, false)

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

	result := SortStrings(string(data), 1, false, true, false, false)

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

	result := SortStrings(string(data), 1, false, false, true, false)

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

	result := SortStrings(string(data), 1, false, false, false, true)

	if result != string(correct) {
		t.Errorf("Incorrect result: numeric sort")
	}
}
