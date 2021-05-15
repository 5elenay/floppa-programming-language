package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ERROR >>> Can't detect input or output argument.")
		os.Exit(1)
	}

	input, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Printf("ERROR >>> Can't open input file (%s).\n", err)
		os.Exit(2)
	}

	inputReaded := string(input)

	flopRegex := regexp.MustCompile(`floppa|caracal|keerr|keeerr|kee?r|foo?o?o?|go|no|flop|hoe`)
	flopKeywords := flopRegex.FindAllString(inputReaded, len(inputReaded))

	compiledProgram := `package main

import "fmt"

func changeByte(mem *[30000]byte, c int, a int, t string) {
	if t == "+" {
		if int((*mem)[c])+a <= 256 {
			(*mem)[c] += byte(a)
		} else {
			(*mem)[c] = byte(a % 255)
		}
	} else {
		if int((*mem)[c])-a >= 0 {
			(*mem)[c] -= byte(a)
		} else {
			(*mem)[c] = byte(256 - (a % 255))
		}
	}
}

func main() {
    memory := [30000]byte{}
    current := 0
    handler := ""
	fmt.Print(handler)


`

	programIndentation := 4

	for _, keyword := range flopKeywords {
		switch keyword {
		case "ker":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 1, \"+\")\n", strings.Repeat(" ", programIndentation))
		case "keer":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 2, \"+\")\n", strings.Repeat(" ", programIndentation))
		case "fo":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 1, \"-\")\n", strings.Repeat(" ", programIndentation))
		case "foo":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 2, \"-\")\n", strings.Repeat(" ", programIndentation))
		case "keerr":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 3, \"+\")\n", strings.Repeat(" ", programIndentation))
		case "keeerr":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 4, \"+\")\n", strings.Repeat(" ", programIndentation))
		case "fooo":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 3, \"-\")\n", strings.Repeat(" ", programIndentation))
		case "foooo":
			compiledProgram += fmt.Sprintf("%schangeByte(&memory, current, 4, \"-\")\n", strings.Repeat(" ", programIndentation))
		case "go":
			compiledProgram += fmt.Sprintf("%scurrent ++\n", strings.Repeat(" ", programIndentation))
		case "no":
			compiledProgram += fmt.Sprintf("%scurrent --\n", strings.Repeat(" ", programIndentation))
		case "flop":
			compiledProgram += fmt.Sprintf("%sfor memory[current] != 0 {\n", strings.Repeat(" ", programIndentation))
			programIndentation += 4
		case "hoe":
			programIndentation -= 4
			compiledProgram += fmt.Sprintf("%s}\n", strings.Repeat(" ", programIndentation))
		case "caracal":
			compiledProgram += fmt.Sprintf("%sfmt.Scanf(\"%%s\", &handler)\n%smemory[current] = handler[0]\n", strings.Repeat(" ", programIndentation), strings.Repeat(" ", programIndentation))
		case "floppa":
			compiledProgram += fmt.Sprintf("%sfmt.Print(string(memory[current]))\n", strings.Repeat(" ", programIndentation))
		}
	}
	compiledProgram += "}"

	writeErr := ioutil.WriteFile(os.Args[2], []byte(compiledProgram), 0666)

	if writeErr != nil {
		fmt.Printf("ERROR >>> Can't write output file (%s).\n", writeErr)
		os.Exit(3)
	}

	fmt.Printf("Compiled Floppa Programming Language Successfully. Check %s!", os.Args[2])
}
