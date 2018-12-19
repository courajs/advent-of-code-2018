package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := make([][]rune, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputs = append(inputs, []rune(s.Text()))
	}
	check(s.Err())

	for _, left := range inputs {
		for _, right := range inputs {
			var diffs int
			var diffAt int
			for i := range left {
				if left[i] != right[i] {
					diffs++
					diffAt = i
				}
			}
			if diffs == 1 {
				fmt.Println(string(left[:diffAt]) + string(left[diffAt+1:]))
				return
			}
		}
	}
	fmt.Println("No find")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
