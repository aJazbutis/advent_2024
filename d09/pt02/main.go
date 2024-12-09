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
	var freeSpace, fileSize int
outer:
	for i < j {
		idx := disc[j]
		j1 := j
		for j >= i && disc[j] == idx {
			j--
		}
		fileSize = j1 - j
		if fileSize == 0 {
			break outer
		}
		for j >= i {
			for j >= i && disc[i] > -1 {
				i++
			}
			i1 := i
			for j >= i1 && disc[i1] == -1 {
				i1++
			}
			freeSpace = i1 - i
			if freeSpace >= fileSize {
				break
			} else {
				i = i1
			}
		}
		if fileSize <= freeSpace {
			for ij := 0; ij < fileSize; ij++ {
				disc[i+ij], disc[j1-ij] = disc[j1-ij], disc[i+ij]
			}
		} else {
			for j > i && disc[j] != -1 {
				j--
			}
		}
		i = 0
		for j > i && disc[i] > -1 {
			i++
		}
		for j > i && disc[j] == -1 {
			j--
		}
	}
	var sum int
	for idx, val := range disc {
		if val < 0 {
			continue
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
