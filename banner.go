package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	ban, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	line := strings.Split(string(ban), "\n")
	result := make(map[rune][]string)
	ch := rune(32)
	for i := 0; i < len(line); i += 9 {
		if i+8 < len(line) {
			result[ch] = line[i+1 : i+9]
			ch++
		}
		if len(result) == 0 {
			return nil, fmt.Errorf("Empty banner file")
		}
	}
	return result, nil
}
