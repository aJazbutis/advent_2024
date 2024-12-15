package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x int
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getGrid() [][]rune {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	grid := [][]rune{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func count(grid [][]rune, visited [][]bool, current Point) int {
	visited[current.y][current.x] = true
	toVisist := []Point{current}
	perimeter, plot, plant := 0, 0, grid[current.y][current.x]
	for len(toVisist) > 0 {
		current = toVisist[0]
		toVisist = toVisist[1:]
		plot++
		for _, direction := range directions {
			x, y := current.x+direction.x, current.y+direction.y
			if x < 0 || y < 0 || y == len(grid) || x == len(grid[0]) || grid[y][x] != plant {
				perimeter++
			} else if !visited[y][x] {
				visited[y][x] = true
				toVisist = append(toVisist, Point{y, x})
			}
		}
	}
	return perimeter * plot
}

func solution() int {
	sum := 0
	grid := getGrid()
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if !visited[y][x] {
				current := Point{y, x}
				sum += count(grid, visited, current)
			}
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	fmt.Println(solution())
}
