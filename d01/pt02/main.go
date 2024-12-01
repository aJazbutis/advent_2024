package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func freq(arr []int) map[int]int {
	ret := make(map[int]int)
	for i := 0; i < len(arr); {
		val := arr[i]
		c := 1
		i++
		for i < len(arr) && arr[i] == val {
			c++
			i++
		}
		ret[val] = c
	}
	return ret
}

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
	freq_l := freq(left)
	freq_r := freq(right)
	sum := 0
	for k, v := range freq_l {
		sum += freq_r[k] * v * k
	}
	fmt.Println(sum)
}
