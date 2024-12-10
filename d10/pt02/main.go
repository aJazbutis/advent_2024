package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getGrid() [][]byte {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	grid := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Bytes())
	}
	return grid
}

type Point struct {
	y, x int
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBounds(y, x, lenY, lenX int) bool {
	return y >= 0 && y < lenY && x >= 0 && x < lenX
}

func getTrailheadScore(g *[][]byte, i, j int) int {
	grid := *g
	score := 0
	visited := make(map[Point]bool)
	toVisit := []Point{{i, j}}
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		visited[current] = true
		for _, direction := range directions {
			y := current.y + direction.y
			x := current.x + direction.x
			if inBounds(y, x, len(grid), len(grid[0])) && !visited[Point{y, x}] {
				if grid[y][x]-grid[current.y][current.x] == 1 {
					if grid[y][x] == '9' {
						score++
					} else {
						toVisit = append(toVisit, Point{y, x})
					}
				}
			}
		}
	}
	return score
}

func sumOfScores() int {
	grid := getGrid()
	sum := 0
	lenY, lenX := len(grid), len(grid[0])
	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			if grid[i][j] == '0' {
				sum += getTrailheadScore(&grid, i, j)
			}
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	fmt.Println(sumOfScores())
}
