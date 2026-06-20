package main

import "strings"

func PrintArt(s string, banner map[rune][]string) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\\n")

	if strings.Join(lines, "") == "" {
		return strings.Repeat("\n", len(lines)-1)
	}

	var result strings.Builder

	for _, inp := range lines {
		if inp == "" {
			result.WriteString("\n")
			continue
		}

		rend := Render(inp, banner)

		for _, k := range rend {
			result.WriteString(k)
			result.WriteString("\n")
		}
	}

	return result.String()
}
