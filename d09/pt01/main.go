package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput() []byte {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Bytes()
}

func spread(diskMap []byte) []int {
	ret := []int{}
	id := -1
	for i := range diskMap {
		size := diskMap[i] - '0'
		if i%2 == 0 {
			id++
			for size > 0 {
				ret = append(ret, id)
				size--
			}
		} else {
			for size > 0 {
				ret = append(ret, -1)
				size--
			}
		}
	}
	return ret
}

func getChecksum(discMap []byte) int {
	disc := spread(discMap)
	i, j := 0, len(disc)-1
outer:
	for {
		for disc[i] > -1 {
			i++
			if i >= j {
				break outer
			}
		}
		for disc[j] == -1 {
			j--
			if i >= j {
				break outer
			}
		}
		if i >= j {
			break outer
		}
		disc[i], disc[j] = disc[j], disc[i]
	}
	sum := 0
	for idx, val := range disc {
		if val < 0 {
			return sum
		}
		sum += idx * val
	}
	return sum
}
func main() {
	if len(os.Args) != 2 {
		log.Fatal("No input")
	}
	diskMap := getInput()
	fmt.Println(getChecksum(diskMap))
}
