package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

const (
	up = iota
	right
	down
	left
)

type Point struct {
	y, x int
}

type Step struct {
	dir, y, x int
}

func isLoop(grid [][]rune, guard Step, directions map[int]Point) bool {
	lenY, lenX := len(grid), len(grid[0])
	path := []Step{}
	for {
		x, y := guard.x, guard.y
		if slices.Contains(path, guard) {
			return true
		}
		path = append(path, guard)
		x += directions[guard.dir].x
		y += directions[guard.dir].y
		if x < 0 || y < 0 || x == lenX || y == lenY {
			return false
		}
		switch grid[y][x] {
		case '#':
			switch guard.dir {
			case up:
				guard.dir = right
			case right:
				guard.dir = down
			case down:
				guard.dir = left
			case left:
				guard.dir = up
			}
		default:
			guard.x = x
			guard.y = y
		}
	}
}

func pathArea(grid [][]rune) int {
	directions := make(map[int]Point)
	short := make(map[rune]int)
	short['v'], short['^'], short['<'], short['>'] = down, up, left, right
	directions[up] = Point{-1, 0}
	directions[down] = Point{1, 0}
	directions[left] = Point{0, -1}
	directions[right] = Point{0, 1}
	ret := 0
	var start, guard Step
	lenY, lenX := len(grid), len(grid[0])
outer:
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			r := grid[y][x]
			switch r {
			case '>', '<', '^', 'v':
				start = Step{short[r], y, x}
				guard = start
				break outer
			}
		}
	}
	path := []Point{}
	for {
		x, y := guard.x, guard.y
		if !slices.Contains(path, Point{guard.y, guard.x}) {
			path = append(path, Point{guard.y, guard.x})
		}
		grid[y][x] = 'X'
		x += directions[guard.dir].x
		y += directions[guard.dir].y
		if x < 0 || y < 0 || x == lenX || y == lenY {
			break
		}
		switch grid[y][x] {
		case '#':
			switch guard.dir {
			case up:
				guard.dir = right
			case right:
				guard.dir = down
			case down:
				guard.dir = left
			case left:
				guard.dir = up
			}
		default:
			guard.x = x
			guard.y = y
		}
	}
	for i := 1; i < len(path); i++ {
		y, x := path[i].y, path[i].x
		grid[y][x] = '#'
		if isLoop(grid, start, directions) {
			ret++
		}
		grid[y][x] = '.'
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("no input")
	}
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	fmt.Println(pathArea(grid))
}
