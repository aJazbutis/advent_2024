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
	c := 0
	lenY := len(s)
	lenX := len(s[0])
	for i := 1; i < lenY -1; i++ {
		for j := 1; j < lenX -1; j++ {
			if s[i][j] == 'A' {
				if ((s[i-1][j-1] == 'M' && s[i+1][j+1] == 'S') || (s[i-1][j-1] == 'S' && s[i+1][j+1] == 'M')) && ((s[i-1][j+1] == 'M' && s[i+1][j-1] == 'S') || (s[i-1][j+1] == 'S' && s[i+1][j-1] == 'M')) {
						c++
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
