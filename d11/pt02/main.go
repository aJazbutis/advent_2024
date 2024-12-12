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

func blink(stones map[string]int) map[string]int {
	s := make(map[string]int)
	for stone, amount := range stones {
		if stone == "0" {
			s["1"] += amount
		} else if len(stone)%2 == 0 {
			twoStones := split(stone)
			s[twoStones[0]] += amount
			s[twoStones[1]] += amount
		} else {
			num, _ := strconv.Atoi(stone)
			num *= 2024
			stone = strconv.Itoa(num)
			s[stone] += amount
		}
	}
	return s
}

func count(stones map[string]int) int {
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}
	sum := 0
	for _, amount := range stones {
		sum += amount
	}
	return sum
}

func getInput() map[string]int {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	stones := make(map[string]int)
	for _, stone := range s {
		stones[stone]++
	}
	return stones
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	stones := getInput()
	fmt.Println(count(stones))
}
