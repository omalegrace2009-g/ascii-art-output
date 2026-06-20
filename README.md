# ASCII Art Output

## Description

ASCII Art Output is a Go program that generates ASCII art from a given string using predefined banner files. In addition to displaying the generated ASCII art in the terminal, the program supports writing the output to a text file using the `--output=<fileName>` flag.

The project extends the basic ASCII-Art functionality by introducing file handling and command-line option parsing.

---

## Author

- Ooja Omale

---

## Usage

### Display ASCII Art in the Terminal

```bash
go run . "hello" standard
```
Save ASCII Art to a File
go run . --output=banner.txt "hello" standard
View Generated File
cat banner.txt


**Implementation Details**

1. Argument Validation

The program validates the number and format of command-line arguments.

Supported formats:

go run . "STRING"
go run . "STRING" BANNER
go run . --output=<fileName> "STRING" BANNER

Any invalid format prints a usage message and exits.

2. Banner Loading

Banner files are loaded using Go’s file system API (os.ReadFile).

Each printable ASCII character (ASCII 32–126) is mapped to its 8-line ASCII representation and stored in:

map[rune][]string

This allows efficient lookup during rendering.

3. Input Validation

All characters in the input string are validated to ensure they are within the printable ASCII range (32–126).

If an invalid character is found, the program stops execution and displays an error message.

4. ASCII Art Rendering

The string is processed character by character.

For each character:

The corresponding 8-line ASCII pattern is retrieved from the banner map
Each line is concatenated horizontally with other characters
This produces the final ASCII-art output

5. File Output 

If the --output= flag is provided:

The filename is extracted from the flag
The generated ASCII art is written using os.WriteFile
No output is printed to the terminal

If no flag is provided, output is printed directly to stdout.

Allowed Packages

Only standard Go packages are used.

Example
go run . --output=banner.txt "hello" standard

Output written to banner.txt:

 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               

---