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

	var contested = make(map[*claim]bool)

	var grid [1000][1000]*claim
	for i, c := range claims {
		for x := c.x; x < c.x+c.width; x++ {
			for y := c.y; y < c.y+c.height; y++ {
				if grid[x][y] == nil {
					grid[x][y] = &claims[i]
				} else {
					contested[&claims[i]] = true
					contested[grid[x][y]] = true
				}
			}
		}
	}

	for i := range claims {
		_, ok := contested[&claims[i]]
		if !ok {
			fmt.Println(claims[i].id)
			return
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
