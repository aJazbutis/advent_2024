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

func checkAndSumMiddle(rules, updates [][]string) int {
	ret := 0
	for _, update := range updates {
		ok := true
		for _, rule := range rules {
			i1 := slices.Index(update, rule[0])
			if i1 != -1 {
				i2 := slices.Index(update, rule[1])
				if i2 != -1 && i1 > i2 {
					ok = false
					break
				}
			}
		}
		if ok {
			val, _ := strconv.Atoi(update[len(update)/2])
			ret += val
		}
	}
	return ret
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
	rules, updates := [][]string{}, [][]string{}
	scanner := bufio.NewScanner(f)
	stillRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			stillRules = false
			continue
		}
		if stillRules {
			rules = append(rules, strings.Split(line, "|"))
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	fmt.Println(checkAndSumMiddle(rules, updates))
}
