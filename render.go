package main

func Render(get string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, rn := range get {
			result[i] += banner[rn][i]
		}
	}
	return result
}
