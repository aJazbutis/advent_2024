package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	y, x int
}

type Robot struct {
	p, v Point
}

func extractPoint(s string) Point {
	ss := strings.Split(s, ",")
	x, _ := strconv.Atoi(ss[0][2:])
	y, _ := strconv.Atoi(ss[1])
	return Point{y, x}
}

func getRobot(s string) Robot {
	ss := strings.Fields(s)
	p, v := extractPoint(ss[0]), extractPoint(ss[1])
	return Robot{p, v}
}

func getRobotList() []Robot {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	robots := []Robot{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		robots = append(robots, getRobot(scanner.Text()))
	}
	return robots
}

func (robot *Robot) Move(time, lenX, lenY int) {
	for time > 0 {
		time--
		nx, ny := robot.p.x+robot.v.x, robot.p.y+robot.v.y
		if nx < 0 {
			nx = lenX + nx
		} else if nx >= lenX {
			nx = nx - lenX
		}
		if ny < 0 {
			ny = lenY + ny
		} else if ny >= lenY {
			ny = ny - lenY
		}
		robot.p.x = nx
		robot.p.y = ny
	}
}

func safetyFactor(robots []Robot, lenX, lenY int) int {
	var q1, q2, q3, q4 int
	for _, robot := range robots {
		if robot.p.y < lenY/2 {
			if robot.p.x < lenX/2 {
				q1++
			} else if robot.p.x > lenX/2 {
				q2++
			}
		} else if robot.p.y > lenY/2 {
			if robot.p.x < lenX/2 {
				q3++
			} else if robot.p.x > lenX/2 {
				q4++
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input")
	}
	lenX, lenY := 101, 103
	time := 100
	robots := getRobotList()
	for i := range robots {
		robots[i].Move(time, lenX, lenY)
	}
	fmt.Println(safetyFactor(robots, lenX, lenY))
}
