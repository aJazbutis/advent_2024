package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split(stone string) []string {
	mid := len(stone) / 2
	stone1 := stone[:mid]
	stone2 := stone[mid:]
	num, _ := strconv.Atoi(stone2)
	stone2 = strconv.Itoa(num)
	return []string{stone1, stone2}
}

func blink(stones []string) []string {
	s := []string{}
	for _, stone := range stones {
		if stone == "0" {
			s = append(s, "1")
		} else if len(stone)%2 == 0 {
			s = append(s, split(stone)...)
		} else {
			num, _ := strconv.Atoi(stone)
			num *= 2024
			stone = strconv.Itoa(num)
			s = append(s, stone)
		}
	}
	return s
}

func count(stones []string) int {
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func getInput() []string {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	stones := strings.Fields(scanner.Text())
	return stones
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	stones := getInput()
	fmt.Println(count(stones))
}
