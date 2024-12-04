package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func xmasCount(arr []string) int {
	s := [][]rune{}
	for _, l := range arr {
		s = append(s, []rune(l))
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, 0},
		{-1, 0},
		{-1, 1},
		{1, -1},
	}
	X := []rune("MAS")
	c := 0
	lenY := len(s)
	lenX := len(s[0])
	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			if s[i][j] == 'X' {
				for _, direction := range directions {
					x, y := j, i
					ok := false
					for idx, r := range X {
						x += direction[0]
						y += direction[1]
						if x >= 0 && x < lenX && y >= 0 && y < lenY {
							if s[y][x] != r {
								ok = false
								break
							} else if idx == 2 {
								ok = true
							}
						}
					}
					if ok {
						c++
					}
				}
			}
		}
	}
	return c
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No input provided.")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(xmasCount(lines))
}
