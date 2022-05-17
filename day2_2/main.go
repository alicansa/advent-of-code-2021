package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

type Position struct {
	depth         int
	horizontalPos int
	aim           int
}

func main() {
	f, err := os.ReadFile("data.txt")

	// log error
	if err != nil {
		log.Fatal(err)
		return
	}

	actions := strings.Split(string(f), "\n")
	var p Position

	// for each action we need the action parameters
	for _, action := range actions {
		parameters := strings.Split(action, " ")
		direction := Direction(parameters[0])
		length, err := strconv.Atoi(parameters[1])

		if err != nil {
			log.Fatal(err)
			return
		}

		calculatePosition(&p, direction, length)
	}

	println(p.depth * p.horizontalPos)
}

func calculatePosition(
	p *Position,
	d Direction,
	length int) {
	switch d {
	case Forward:
		p.horizontalPos += length
		p.depth += p.aim * length
		break
	case Down:
		p.aim += length
		break
	case Up:
		if p.aim <= length {
			p.aim = 0
			p.depth = 0
			break
		}

		p.aim -= length
		break
	}
}
