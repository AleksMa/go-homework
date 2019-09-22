package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	keyFlag         = flag.Int("k", 1, "an int")
	ignoreCaseFlag  = flag.Bool("f", false, "a bool")
	uniqueFlag      = flag.Bool("u", false, "a bool")
	reverseFlag     = flag.Bool("r", false, "a bool")
	outputFlag      = flag.String("o", "", "a string")
	numericSortFlag = flag.Bool("n", false, "a bool")
)

const (
	notNumericType = iota
	intType
	floatType
)

type Options struct {
	keyFlag         int
	ignoreCaseFlag  bool
	uniqueFlag      bool
	reverseFlag     bool
	outputFlag      string
	numericSortFlag bool
}

type Line struct {
	Words    []string
	PureLine string
}

func (l Line) String() string {
	return fmt.Sprintf("%s", l.PureLine)
}

func WordsLess(x []string, y []string, numericFlag bool) bool {
	if len(x) == 0 {
		return true
	} else if len(y) == 0 {
		return false
	} else if x[0] == y[0] {
		return WordsLess(x[1:], y[1:], numericFlag)
	} else if x[0] == "" && !numericFlag {
		return true
	} else if y[0] == "" && !numericFlag {
		return false
	} else if numericFlag {
		var num1, num2 float64
		status1, status2 := notNumericType, notNumericType

		if tempNum, err := strconv.ParseFloat(x[0], 64); err == nil {
			num1 = tempNum
			status1 = floatType
		} else if tempNum, err := strconv.Atoi(x[0]); err == nil {
			num1 = float64(tempNum)
			status1 = intType
		}

		if tempNum, err := strconv.ParseFloat(y[0], 64); err == nil {
			num2 = tempNum
			status2 = floatType
		} else if tempNum, err := strconv.Atoi(y[0]); err == nil {
			num2 = float64(tempNum)
			status2 = intType
		}

		switch {
		case status1 != notNumericType && status2 != notNumericType:
			return num1 < num2
		case status1 != notNumericType:
			return num1 < 0
		case status2 != notNumericType:
			return 0 < num2
		}
	}
	return x[0] < y[0]
}

func SortStrings(data string, options Options) string {
	uniqueSet := make(map[string]bool)

	linesBuf := strings.Split(data, "\n")
	lines := make([]Line, 0)

	for _, line := range linesBuf {

		if options.uniqueFlag {
			uniqueLine := line
			if options.ignoreCaseFlag {
				uniqueLine = strings.ToLower(uniqueLine)
			}
			if uniqueSet[uniqueLine] {
				continue
			} else {
				uniqueSet[uniqueLine] = true
			}
		}

		j := len(lines)
		lines = append(lines, Line{})

		words := strings.Split(line, " ")
		if options.ignoreCaseFlag {
			words = strings.Split(strings.ToLower(line), " ")
		}
		wordsLen := len(words)
		minKey := int(math.Min(float64(options.keyFlag-1), float64(wordsLen)))
		wordsBase := words[minKey:]
		wordsTail := words[:minKey]
		lines[j].Words = make([]string, len(wordsBase), len(wordsBase))
		copy(lines[j].Words, wordsBase)
		lines[j].Words = append(lines[j].Words, wordsTail...)
		lines[j].PureLine = line
	}

	if options.reverseFlag {
		sort.Slice(lines, func(i, j int) bool {
			return !WordsLess(lines[i].Words, lines[j].Words, options.numericSortFlag)
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			return WordsLess(lines[i].Words, lines[j].Words, options.numericSortFlag)
		})
	}

	var result strings.Builder

	for i, line := range lines {
		result.WriteString(line.PureLine)
		if i < len(lines)-1 || options.outputFlag != "" { // Unix sort util prints '\n' at end only in -o mode
			result.WriteString("\n")
		}
	}
	return result.String()
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: \"go run sort.go [-k int] [-f] [-u] [-r] [-o filename] [-n] filename\"")
		return
	}
	path := flag.Args()[0]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	options := Options{
		keyFlag:         *keyFlag,
		ignoreCaseFlag:  *ignoreCaseFlag,
		uniqueFlag:      *uniqueFlag,
		reverseFlag:     *reverseFlag,
		outputFlag:      *outputFlag,
		numericSortFlag: *numericSortFlag,
	}

	result := SortStrings(string(data), options)

	if *outputFlag == "" {
		fmt.Println(result)
	} else {
		err := ioutil.WriteFile(*outputFlag, []byte(result), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
