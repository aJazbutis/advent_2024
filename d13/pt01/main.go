package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Step struct {
	Point
	cost int
}

type State struct {
	x, y, cost, c1, c2 int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq) - 1
	ret := (*pq)[n]
	*pq = (*pq)[:n]
	return ret
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minCost(a, b Step, prize Point) (int, bool) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{0, 0, 0, 0, 0})
	visited := make(map[[4]int]int)
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(State)
		x, y, cost, c1, c2 := curr.x, curr.y, curr.cost, curr.c1, curr.c2
		if x == prize.x && y == prize.y {
			return cost, true
		}
		stateKey := [4]int{x, y, c1, c2}
		if v, ok := visited[stateKey]; ok && v <= cost {
			continue
		}
		visited[stateKey] = cost
		if c1 < 100 {
			nx, ny := x+a.x, y+a.y
			heap.Push(pq, State{nx, ny, cost + a.cost, c1 + 1, c2})
		}
		if c2 < 100 {
			nx, ny := x+b.x, y+b.y
			heap.Push(pq, State{nx, ny, cost + b.cost, c1, c2 + 1})
		}
	}
	return -1, false
}

func parse(ss []string) Point {
	x, _ := strconv.Atoi(ss[0][2 : len(ss[0])-1])
	y, _ := strconv.Atoi(ss[1][2:])
	return Point{y, x}
}

func tokens(lines []string) int {
	ss := strings.Fields(lines[0])
	p := parse(ss[2:])
	a := Step{p, 3}
	ss = strings.Fields(lines[1])
	p = parse(ss[2:])
	b := Step{p, 1}
	ss = strings.Fields(lines[2])
	prize := parse(ss[1:])
	// fmt.Println(a, b, prize)
	gcdX, gcdY := gcd(a.x, b.x), gcd(a.y, b.y)
	if !(prize.x%gcdX == 0 && prize.y%gcdY == 0) {
		return 0
	}
	t, ok := minCost(a, b, prize)
	if ok {
		return t
	}
	return 0
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	lines := []string{}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
		if len(lines) == 3 {
			sum += tokens(lines)
			lines = []string{}
		}
	}
	fmt.Println(sum)
}
