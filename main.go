package main

import (
	"os"
	"bufio"
	"time"
	"fmt"
	"flag"
	"strings"
)

func readFile(path string) []string {
	var result []string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

var d bool
func initFlags() {
	flag.BoolVar(&d, "d", false, "")
	flag.Parse()
}

func hasThreeVowels(name string) bool {
	vowels := "aeuoi"
	splittedName := strings.Split(name, "")
	vowelsCount := 0

	for _, char := range(splittedName) {
		if strings.Contains(vowels, char) {
			vowelsCount++
		}
	}
	return vowelsCount > 2
}

func hasWrongSubString(name string) bool {
	wrongStrings := []string{"ab", "cd", "pq", "xy"}

	for _, str := range(wrongStrings) {
		if strings.Contains(name, str) {
			return true
		}
	}

	return false
}

func hasDuplicateChar(name string) bool {
	chars := strings.Split(name, "")
	charsLen := len(chars)
	for i := 1; i < charsLen; i++ {
		if chars[i] == chars[i-1] {
			return true
		}
	}
	return false
}

func simple() {
	names := readFile("input.txt")
	goodNamesCount := 0
	for _, name := range(names) {
		if !hasWrongSubString(name) && hasThreeVowels(name) && hasDuplicateChar(name) {
			goodNamesCount++
		}
	}
	fmt.Println("Result = ", goodNamesCount)
}

func hasSameLetterAfterLetter(name string) bool {
	chars := strings.Split(name, "")
	charsLen := len(chars)
	for i := 2; i < charsLen; i++ {
		if chars[i] == chars[i-2] {
			return true
		}
	}
	return false
}

func hasDuplicatePare(name string) bool {
	nameLen := len(name)

	for i := 2; i < nameLen; i++ {
		pare := name[i-2:i]
		if strings.Count(name, pare) > 1 {
			return true
		}
	}

	return false
}

func difficult() {
	names := readFile("input.txt")
	goodNamesCount := 0
	for _, name := range(names) {
		if hasSameLetterAfterLetter(name) && hasDuplicatePare(name) {
			goodNamesCount++
		}
	}
	fmt.Println("Result = ", goodNamesCount)
}

func main() {
	start := time.Now()
	initFlags()
	if d {
		difficult()
	} else {
		simple()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
