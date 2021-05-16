package main

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


    changeByte(&memory, current, 4, "+")
    changeByte(&memory, current, 4, "+")
    changeByte(&memory, current, 4, "+")
    changeByte(&memory, current, 1, "+")
    for memory[current] != 0 {
        current ++
        changeByte(&memory, current, 4, "+")
        changeByte(&memory, current, 1, "+")
        current --
        changeByte(&memory, current, 1, "-")
    }
    current ++
    fmt.Print(string(memory[current]))
    if memory[current] == byte(65) {
        changeByte(&memory, current, 1, "+")
    }
    fmt.Print(string(memory[current]))
}