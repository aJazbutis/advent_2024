package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getEquations(path string) [][]int {
	equations := [][]int{}
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		equation := []int{}
		line := scanner.Text()
		lin := strings.Split(line, ":")
		result, _ := strconv.Atoi(lin[0])
		equation = append(equation, result)
		li := strings.Fields(lin[1])
		for _, l := range li {
			num, _ := strconv.Atoi(l)
			equation = append(equation, num)
		}
		equations = append(equations, equation)
	}
	return equations
}

func process(nums []int, target, current, idx int) bool {
	if idx == len(nums) {
		return current == target
	}
	return process(nums, target, current+nums[idx], idx+1) || process(nums, target, current*nums[idx], idx+1)
}

func canSolve(nums []int) bool {
	return process(nums[1:], nums[0], nums[1], 1)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No input")
	}
	sum := 0
	equations := getEquations(os.Args[1])
	for _, equation := range equations {
		if canSolve(equation) {
			sum += equation[0]
		}
	}
	fmt.Println(sum)
}
