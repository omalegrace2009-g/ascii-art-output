package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func main() {
	args := os.Args

	if len(args) < 2 || len(args) > 4 {
		usage()
		return
	}

	var (
		option string
		str    string
		banner string = "standard"
	)

	switch len(args) {

	// go run . "hello"
	case 2:
		str = args[1]

	// go run . "hello" shadow
	case 3:
		str = args[1]
		banner = args[2]

	// go run . --output=file.txt "hello" shadow
	case 4:
		option = args[1]
		str = args[2]
		banner = args[3]

		if !strings.HasPrefix(option, "--output=") || len(option) <= len("--output=") {
			usage()
			return
		}
	}

	// validate input string
	if _, err := Validate(str); err != nil {
		fmt.Printf("Error: invalid character '%c'\n", err)
		return
	}

	file := "banners/" + banner + ".txt"
	bannerMap, err := LoadBanner(file)
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}

	output := PrintArt(str, bannerMap)

	if option != "" {
		filename := strings.TrimPrefix(option, "--output=")
		if filename == "" {
			usage()
			return
		}

		err := os.WriteFile(filename, []byte(output), 0644) 
		if  err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Done!!")
		return
	}

	fmt.Print(output)
}
