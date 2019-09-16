package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	//"sort"
	"strings"
)

type Line struct {
	Words    []string
	PureLine string
}

func (l Line) String() string {
	return fmt.Sprintf("%s", l.PureLine)
}

func WordsLess(x []string, y []string) bool {
	if len(x) == 0 {
		return true
	} else if len(y) == 0 {
		return false
	} else if x[0] == y[0] {
		return WordsLess(x[1:], y[1:])
	} else {
		return x[0] < y[0]
	}
}

type ByWords []Line

func (ls ByWords) Len() int      { return len(ls) }
func (ls ByWords) Swap(i, j int) { ls[i], ls[j] = ls[j], ls[i] }
func (ls ByWords) Less(i, j int) bool {
	return WordsLess(ls[i].Words, ls[j].Words)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: \"go run sort.go filename\"")
		return
	}

	// map[string]bool

	path := os.Args[1]

	data, _ := ioutil.ReadFile(path)
	linesBuf := strings.Split(string(data), "\n")
	lines := make([]Line, len(linesBuf))

	for i, line := range linesBuf {
		words := strings.Split(line, " ")
		lines[i].Words = make([]string, len(words), len(words))
		copy(lines[i].Words, words)
		lines[i].PureLine = line
	}

	sort.Sort(ByWords(lines))

	for _, line := range lines {
		fmt.Printf("%s\n", line.PureLine)
	}

}
