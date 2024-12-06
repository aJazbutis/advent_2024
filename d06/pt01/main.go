package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func pathArea(grid [][]rune) int {
	directions := make(map[int]Point)
	short := make(map[rune]int)
	short['v'], short['^'], short['<'], short['>'] = down, up, left, right
	directions[up] = Point{-1, 0}
	directions[down] = Point{1, 0}
	directions[left] = Point{0, -1}
	directions[right] = Point{0, 1}
	ret := 0
	var dir int
	var guard Point
	lenY, lenX := len(grid), len(grid[0])
outer:
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			r := grid[y][x]
			switch r {
			case '>', '<', '^', 'v':
				dir = short[r]
				guard = Point{y, x}
				break outer
			}
		}
	}
	for {
		x, y := guard.x, guard.y
		grid[y][x] = 'X'
		x += directions[dir].x
		y += directions[dir].y
		if x < 0 || y < 0 || x == lenX || y == lenY {
			break
		}
		switch grid[y][x] {
		case '#':
			switch dir {
			case up:
				dir = right
			case right:
				dir = down
			case down:
				dir = left
			case left:
				dir = up
			}
		default:
			guard.x = x
			guard.y = y
		}
	}
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			switch grid[y][x] {
			case 'X':
				ret++
			}
		}
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
