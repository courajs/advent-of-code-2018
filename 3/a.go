package main

import (
	"bufio"
	"fmt"
	"os"
)

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func main() {
	claims := make([]claim, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		c := claim{}
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.width, &c.height)
		claims = append(claims, c)
		check(err)
	}
	check(s.Err())

	var grid [1000][1000]int8
	for _, c := range claims {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if grid[i][j] < 2 {
					grid[i][j]++
				}
			}
		}
	}

	var sum int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				sum++
			}
		}
	}

	fmt.Println(sum)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
