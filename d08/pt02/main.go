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

func getAntennaList(grid [][]byte) map[byte][]Point {
	antennas := make(map[byte][]Point)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tile := grid[i][j]
			if tile != '.' {
				_, exists := antennas[tile]
				if exists {
					antennas[tile] = append(antennas[tile], Point{i, j})
				} else {
					antennas[tile] = []Point{{i, j}}
				}
			}
		}
	}
	return antennas
}

func getAntinodes(antinodes *map[Point]struct{}, a1, a2 Point, lenY, lenX int) {
	dist := Point{a2.y - a1.y, a2.x - a1.x}
	x := a1.x - dist.x
	y := a1.y - dist.y
	for x >= 0 && y >= 0 && x < lenX && y < lenY {
		(*antinodes)[Point{y, x}] = struct{}{}
		x -= dist.x
		y -= dist.y
	}
	x = a2.x + dist.x
	y = a2.y + dist.y
	for x >= 0 && y >= 0 && x < lenX && y < lenY {
		(*antinodes)[Point{y, x}] = struct{}{}
		x += dist.x
		y += dist.y
	}
}

func antinodeCount(grid [][]byte) int {
	antennas := getAntennaList(grid)
	antinodes := make(map[Point]struct{})
	for _, antennaType := range antennas {
		for i := 0; i < len(antennaType)-1; i++ {
			antenna := antennaType[i]
			for _, secondAntenna := range antennaType[i+1:] {
				getAntinodes(&antinodes, antenna, secondAntenna, len(grid), len(grid[0]))
				antinodes[antenna] = struct{}{}
				antinodes[secondAntenna] = struct{}{}
			}
		}
	}
	return len(antinodes)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No input")
	}
	grid := getGrid()
	fmt.Println(antinodeCount(grid))
}
