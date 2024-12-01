package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	left := make(map[int]int)
	right := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		nums := strings.Fields(line)
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		_, exist := left[l]
		if exist {
			left[l]++
		} else {
			left[l] = 1
		}
		r, err := strconv.Atoi(nums[1])
		_, exist = right[r]
		if exist {
			right[r]++
		} else {
			right[r] = 1
		}
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	sum := 0
	for k, v := range left {
		sum += right[k] * v * k
	}
	fmt.Println(sum)
}
