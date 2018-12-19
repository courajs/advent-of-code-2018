package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := make([]string, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputs = append(inputs, s.Text())
	}
	check(s.Err())

	var twos, threes int
	for _, v := range inputs {
		counts := make(map[rune]int)
		for _, c := range v {
			counts[c]++
		}
		for _, count := range counts {
			if count == 2 {
				twos++
				break
			}
		}
		for _, count := range counts {
			if count == 3 {
				threes++
				break
			}
		}
	}

	fmt.Println(twos * threes)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
