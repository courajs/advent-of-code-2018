package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var freq int64
	var seen = make(map[int64]bool)
	var inputs = make([]int64, 0)
	seen[0] = true
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		n, err := strconv.ParseInt(s.Text(), 0, 0)
		check(err)
		inputs = append(inputs, n)
		freq += n
		_, ok := seen[freq]
		if ok {
			fmt.Println(freq)
			return
		}
		seen[freq] = true
	}
	check(s.Err())
	for {
		for _, n := range inputs {
			freq += n
			_, ok := seen[freq]
			if ok {
				fmt.Println(freq)
				return
			}
			seen[freq] = true
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
