package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int64
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		n, err := strconv.ParseInt(s.Text(), 0, 0)
		check(err)
		sum += n
	}
	check(s.Err())
	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
