package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	isSorted := sort.SliceIsSorted(report, func(i, j int) bool { return report[i] < report[j] })
	if !isSorted {
		isSorted = sort.SliceIsSorted(report, func(i, j int) bool { return report[i] > report[j] })
	}
	if !isSorted {
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		diff = int(math.Abs(float64(diff)))
		if diff < 1 || diff > 3 {
			return false
		}

	}
	return true
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
