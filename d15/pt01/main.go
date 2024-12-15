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

func parse() ([][]rune, []rune) {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := [][]rune{}
	instructions := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			grid = append(grid, []rune(line))
		} else {
			instructions += line
		}
	}
	return grid, []rune(instructions)
}

func findRobot(grid [][]rune) Point {
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if grid[y][x] == '@' {
				return Point{y, x}
			}
		}
	}
	return Point{-1, -1}
}

func canMove(robot Point, grid [][]rune, direction Point) (bool, Point) {
	x, y := robot.x+direction.x, robot.y+direction.y
	for grid[y][x] != '#' {
		if grid[y][x] == '.' {
			return true, Point{y, x}
		} else {
			x += direction.x
			y += direction.y
		}
	}
	return false, Point{}
}

func move(robot Point, grid [][]rune, direction Point) Point {
	ok, freeSpot := canMove(robot, grid, direction)
	if ok {
		for freeSpot != robot {
			prev := Point{freeSpot.y - direction.y, freeSpot.x - direction.x}
			grid[prev.y][prev.x], grid[freeSpot.y][freeSpot.x] = grid[freeSpot.y][freeSpot.x], grid[prev.y][prev.x]
			freeSpot = prev
		}
		return Point{robot.y + direction.y, robot.x + direction.x}
	}
	return robot
}

func followInstructions(grid [][]rune, instructions []rune) {
	directions := make(map[rune]Point)
	directions['v'] = Point{1, 0}
	directions['<'] = Point{0, -1}
	directions['>'] = Point{0, 1}
	directions['^'] = Point{-1, 0}
	robot := findRobot(grid)
	if robot.x == -1 {
		log.Fatal("No robot")
	}
	grid[robot.y][robot.x] = '.'
	for _, instruction := range instructions {
		robot = move(robot, grid, directions[instruction])
	}
}

func gps(p Point) int {
	return p.y*100 + p.x
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	grid, instuctions := parse()
	followInstructions(grid, instuctions)
	gpsSum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				gpsSum += gps(Point{y, x})
			}
		}
	}
	fmt.Println(gpsSum)
}
