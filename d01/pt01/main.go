package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No input prvided.")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	left := []int{}
	right := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		nums := strings.Fields(line)
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		left = append(left, l)
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	fmt.Println(sum)
}
