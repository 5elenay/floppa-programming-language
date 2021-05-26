package main

import "fmt"
import "os"
import "bufio"

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

	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 3, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	changeByte(&memory, current, 3, "+")
	fmt.Print(string(memory[current]))
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	changeByte(&memory, current, 3, "-")
	fmt.Print(string(memory[current]))
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 1, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 1, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	fmt.Print(string(memory[current]))
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 1, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	fmt.Print(string(memory[current]))
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 3, "+")
	fmt.Print(string(memory[current]))
	current++
	changeByte(&memory, current, 4, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 3, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	fmt.Print(string(memory[current]))
	current++
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	memory[current] = str[0]
	current++
	str, _ = reader.ReadString('\n')
	memory[current] = str[0]
	current--
	for memory[current] != 0 {
		current++
		current++
		changeByte(&memory, current, 1, "+")
		current--
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 1, "-")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	current++
	changeByte(&memory, current, 4, "+")
	changeByte(&memory, current, 1, "+")
	changeByte(&memory, current, 1, "+")
	for memory[current] != 0 {
		current++
		changeByte(&memory, current, 4, "+")
		changeByte(&memory, current, 1, "+")
		changeByte(&memory, current, 3, "+")
		current--
		changeByte(&memory, current, 1, "-")
	}
	current--
	for memory[current] != 0 {
		current++
		current++
		changeByte(&memory, current, 1, "+")
		current--
		current--
		changeByte(&memory, current, 1, "-")
	}
	current++
	current++
	fmt.Print(string(memory[current]))
}
