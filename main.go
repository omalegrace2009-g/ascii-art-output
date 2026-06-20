package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX Usage: go run . --output=<fileName.txt> something standard")
		return
	}

	if len(os.Args) == 3 {
		_, err := Validate(os.Args[1])
		if err != nil {
			fmt.Printf("character %v is not an ascii character", os.Args[1])
			return
		}

		file := "banners/" + os.Args[2] + ".txt"
		banner, err := LoadBanner(file)
		if err != nil {
			fmt.Println("Error Reading file", err)
			return
		}
		gen := PrintArt(os.Args[1], banner)
		fmt.Print(gen)
		return
	}

	if len(os.Args) == 4 {
		if !strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}

		file := strings.TrimPrefix(os.Args[1], "--output=")
		if file == "" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}
		_, err := Validate(os.Args[2])
		if err != nil {
			fmt.Printf("character %v is not an ascii character", os.Args[2])
			return
		}

		files := "banners/" + os.Args[3] + ".txt"
		banner, err := LoadBanner(files)
		if err != nil {
			fmt.Println("Error Reading file", err)
			return
		}

		gen := PrintArt(os.Args[2], banner)
		err = os.WriteFile(file, []byte(gen), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Done!")
		return
	}
}
