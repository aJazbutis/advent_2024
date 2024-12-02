package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeAsc(s []int) bool {
	for i := 1; i < len(s)-1; i++ {
		diff1 := s[i] - s[i-1]
		diff2 := s[i+1] - s[i]
		if diff1 < 1 || diff1 > 3 || diff2 > 3 || diff2 < 1 {
			return false
		}
	}
	return true
}

func isSafeDesc(s []int) bool {
	for i := 1; i < len(s)-1; i++ {
		diff1 := s[i-1] - s[i]
		diff2 := s[i] - s[i+1]
		if diff1 < 1 || diff1 > 3 || diff2 > 3 || diff2 < 1 {
			return false
		}
	}
	return true
}

func isVerySafe(s []int) bool {
	safe := isSafeAsc(s)
	if safe {
		return safe
	}
	return isSafeDesc(s)
}

func isSafe(report []int) bool {
	if isVerySafe(report) {
		return true
	} else {
		for i := 0; i < len(report); i++ {
			damp := append([]int{}, report[:i]...)
			damp = append(damp, report[i+1:]...)
			if isVerySafe(damp) {
				return true
			}
		}
	}
	return false
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		nums := strings.Fields(line)
		levels := []int{}
		for _, num := range nums {
			level, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}
		if isSafe(levels) {
			sum++
		}
	}
	fmt.Println(sum)
}
