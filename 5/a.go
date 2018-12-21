package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var chars [50001]rune
	var stack_head = 0

	var char rune
	var err error
	for err == nil {
		char, _, err = r.ReadRune()
		if err == io.EOF {
			break
		} else {
			check(err)
		}
		chars[stack_head] = char
		stack_head++

		for stack_head >= 2 && Reacts(chars[stack_head-1], chars[stack_head-2]) {
			stack_head -= 2
		}
	}

	fmt.Println(stack_head - 1)
}

func Reacts(a, b rune) bool {
	return a != b && unicode.ToLower(a) == unicode.ToLower(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
