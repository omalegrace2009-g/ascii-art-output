package main

import "fmt"

func Validate(s string) (rune, error) {
	for _, rn := range s {
		if rn < 32 || rn > 126 {
			return rn, fmt.Errorf("Invalid Character %c", rn)
		}
	}
	return 0, nil
}
