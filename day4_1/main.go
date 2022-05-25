package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type boardTile = struct {
	isMarked bool
	value    string
}

func main() {
	//get inputs
	f, err := os.ReadFile("data.txt")

	// log error
	if err != nil {
		log.Fatal(err)
		return
	}

	inputs := strings.Split(string(f), "\n\n")

	// the first input should be the numbers to be drawn
	// numbers := inputs[0]

	// rest of the inputs are the boards
	// build boards
	numOfBoards := len(inputs) - 1
	boards := buildBoards(inputs[1:], numOfBoards)

	//loop over inputs and mark boards
	drawnNumbers := strings.Split(inputs[0], ",")
	for _, drawnNumber := range drawnNumbers {
		for _, board := range boards {
			markBoard(board, drawnNumber)
		}

		// on each loop check if there is a winner
		winner := getWinner(boards)

		if winner != nil {
			// calculate winner points
			sumOfUnmarkedTiles, err := sumUnmarkedTiles(winner)
			parsedValue, err := strconv.Atoi(drawnNumber)
			// log error
			if err != nil {
				log.Fatal(err)
				return
			}

			println(sumOfUnmarkedTiles * parsedValue)
			break
		}

	}
}

func sumUnmarkedTiles(board [][]boardTile) (int, error) {
	var sum int
	for _, boardRow := range board {
		for _, boardTile := range boardRow {
			if boardTile.isMarked {
				continue
			}
			parsedValue, err := strconv.Atoi(boardTile.value)

			if err != nil {
				return sum, err
			}

			sum += parsedValue
		}
	}

	return sum, nil
}

func getWinner(boards [][][]boardTile) [][]boardTile {
	// marked values should fill a column or row in the board
	for _, board := range boards {
		var markedRows [5]int
		var markedColumns [5]int
		for y, boardRow := range board {
			for x, boardTile := range boardRow {

				if !boardTile.isMarked {
					continue
				}

				markedRows[y]++
				markedColumns[x]++

				// if any marked rows or columns have 5 then thats a winner
				if markedRows[y] >= 5 || markedColumns[x] >= 5 {
					return board
				}
			}
		}
	}

	return nil
}

func markBoard(board [][]boardTile, value string) {
	for _, boardRow := range board {
		for i, _ := range boardRow {
			if boardRow[i].isMarked || boardRow[i].value != value {
				continue
			}

			boardRow[i].isMarked = true
			return
		}
	}
}

func buildBoards(boardInputs []string, numOfBoards int) [][][]boardTile {

	boards := make([][][]boardTile, numOfBoards)

	for i, boardInput := range boardInputs {
		boardInputLines := strings.Split(boardInput, "\n")
		boards[i] = make([][]boardTile, 5)
		for y, boardInputLine := range boardInputLines {

			boardInputTiles := strings.Split(boardInputLine, " ")

			for _, boardInputTile := range boardInputTiles {

				// one digit values have an space
				// in front of them so we must get rid of empty values
				if boardInputTile == " " || boardInputTile == "" {
					continue
				}

				boards[i][y] = append(boards[i][y], boardTile{
					value:    boardInputTile,
					isMarked: false,
				})
			}
		}
	}

	return boards
}
