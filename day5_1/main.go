package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate = struct {
	x int
	y int
}

type line = struct {
	start coordinate
	end   coordinate
}

func main() {
	// read all the inputs

	//get inputs
	f, err := os.ReadFile("data.txt")

	// log error
	if err != nil {
		log.Fatal(err)
		return
	}

	// get the line inputs
	lineInputs := strings.Split(string(f), "\n")

	// create initial 1x1 diagram
	diagram := make([][]int, 1)
	diagram[0] = make([]int, 1)

	// parse the line inputs
	for _, lineInput := range lineInputs {
		coordinateInputs := strings.Split(lineInput, " -> ")
		startCoordinateInputs := strings.Split(coordinateInputs[0], ",")

		startX, err := strconv.Atoi(startCoordinateInputs[0])
		startY, err := strconv.Atoi(startCoordinateInputs[1])

		endCoordinateInputs := strings.Split(coordinateInputs[1], ",")
		endX, err := strconv.Atoi(endCoordinateInputs[0])
		endY, err := strconv.Atoi(endCoordinateInputs[1])

		// log error
		if err != nil {
			log.Fatal(err)
			return
		}

		diagram = markDiagram(diagram, line{
			start: coordinate{
				x: startX,
				y: startY,
			},
			end: coordinate{
				x: endX,
				y: endY,
			},
		})
	}

	println(countOverlapsEqualOrGreaterThanTwo(diagram))
	// we don't know how big the board should be
	// we can append as we go?
}

func countOverlapsEqualOrGreaterThanTwo(diagram [][]int) int {
	var counter int
	for _, row := range diagram {
		for _, cell := range row {
			if cell >= 2 {
				counter++
			}
		}
	}

	return counter
}

func markDiagram(diagram [][]int, line line) [][]int {

	// only consider vertical or horizontal lines
	if line.start.x != line.end.x && line.start.y != line.end.y {
		return diagram
	}

	//resize the diagram if necessary
	startX := line.start.x
	endX := line.end.x

	if line.end.x < line.start.x {
		endX = line.start.x
		startX = line.end.x
	}

	startY := line.start.y
	endY := line.end.y

	if line.end.y < line.start.y {
		endY = line.start.y
		startY = line.end.y
	}

	if len(diagram) <= endY {
		// needs to add Y - len diagram times current x
		size := endY - len(diagram) + 1
		for i := 0; i < size; i++ {
			newRows := make([]int, len(diagram[0]))
			diagram = append(diagram, newRows)
		}
	}

	if len(diagram[0]) <= endX {
		// we need to loop over each row
		for i := range diagram {
			newColumns := make([]int, endX-len(diagram[i])+1)
			diagram[i] = append(diagram[i], newColumns...)
		}
	}

	// mark diagram

	if startX == endX {
		// vertical change
		for i := startY; i <= endY; i++ {
			diagram[i][startX]++
		}
		return diagram
	}

	// horizontal change
	for i := startX; i <= endX; i++ {
		diagram[startY][i]++
	}
	return diagram
}
