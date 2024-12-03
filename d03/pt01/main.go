package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func extractIntS(s string) string {
	retS := ""
	for i := 0; unicode.IsDigit(rune(s[i])) && i < len(s); i++ {
		retS += string(rune(s[i]))
	}
	return retS
}

func mul(s string) int {
	needle := "mul("
	needleLen := len(needle)
	idx := strings.Index(s, needle)
	end := len(s) - 1
	num := 0
	for idx >= 0 && idx < end {
		idx += needleLen
		x := extractIntS(s[idx:])
		i := len(x)
		if i != 0 && s[idx+i] == ',' {
			idx += i + 1
			y := extractIntS(s[idx:])
			i = len(y)
			if i != 0 && s[idx+i] == ')' {
				idx += i + 1
				X, _ := strconv.Atoi(x)
				Y, _ := strconv.Atoi(y)
				num += X * Y
			}
		}
		s = s[idx:]
		end = len(s) - 1
		idx = strings.Index(s, needle)
	}
	return num
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
		sum += mul(line)
	}
	fmt.Println(sum)
}
