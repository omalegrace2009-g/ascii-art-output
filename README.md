# ASCII Art Output

## Description

ASCII Art Output is a Go program that generates ASCII art from a given string using predefined banner files. In addition to displaying the generated ASCII art on the terminal, the program supports writing the output to a text file using the `--output=<fileName>` flag.

The project extends the functionality of the ASCII-Art project by introducing file handling and command-line option parsing.

---

## Author

* Ooja Omale


---

## Usage

### Display ASCII Art in the Terminal

```bash
go run . "hello" standard
```

### Save ASCII Art to a File

```bash
go run . --output=banner.txt "hello" standard
```

### View the Generated File

```bash
cat banner.txt
```

---

## Implementation Details

### 1. Argument Validation

The program validates the number and format of command-line arguments.

Accepted formats:

```bash
go run . "STRING" BANNER
```

```bash
go run . --output=<fileName.txt> "STRING" BANNER
```

Invalid formats display the usage message.

---

### 2. Banner Loading

The selected banner file is read using Go's file system API.

Each printable ASCII character (ASCII 32–126) is mapped to its corresponding 8-line ASCII-art representation and stored in a map:

```go
map[rune][]string
```

This allows constant-time lookup during rendering.

---

### 3. Input Validation

Before rendering, the program verifies that all characters belong to the printable ASCII range (32–126).

Invalid characters cause the program to terminate with an error message.

---

### 4. ASCII Art Rendering

The input string is processed character by character.

For each character:

1. The corresponding ASCII-art pattern is retrieved from the banner map.
2. The correct line of the pattern is appended to the output.
3. The process is repeated for all 8 lines of the banner.

This constructs the final ASCII-art representation.

---

### 5. File Output

When the `--output=` flag is provided:

1. The filename is extracted from the flag.
2. The generated ASCII art is written to the specified file using:

```go
os.WriteFile()
```

3. No output is printed to the terminal.

Without the flag, the ASCII art is displayed directly in the terminal.

---

## Allowed Packages

Only Go standard library packages are used.

---

## Example

Command:

```bash
go run . --output=banner.txt "hello" standard
```

Output saved in:

```text
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
```
