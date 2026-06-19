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

	for _, g := range lines {
		if g == "" {
			result.WriteString("\n")
			continue
		}

		rd := Render(g, banner)

		for _, k := range rd {
			result.WriteString(k)
			result.WriteString("\n")
		}
	}

	return result.String()
}
