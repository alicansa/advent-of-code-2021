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

type board = struct {
	isWinner bool
	value    [][]boardTile
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

	// rest of the inputs are the boards
	// build boards
	numOfBoards := len(inputs) - 1
	boards := buildBoards(inputs[1:], numOfBoards)

	//loop over inputs and mark boards
	drawnNumbers := strings.Split(inputs[0], ",")
	var lastWinner board
	for k, drawnNumber := range drawnNumbers {

		for i := range boards {

			if boards[i].isWinner {
				continue
			}

			markBoard(boards[i].value, drawnNumber)
		}

		// on each loop check if there is a winner
		winners := getWinners(boards)
		numOfWinners := getNumberOfWinners(boards)

		if len(winners) > 0 {
			lastWinner = winners[0]
		}

		if k == len(drawnNumbers)-1 || numOfWinners == len(boards) {
			sumOfUnmarkedTiles, err := sumUnmarkedTiles(lastWinner.value)
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

func getNumberOfWinners(boards []board) int {
	var numOfWinners int
	for _, board := range boards {
		if board.isWinner {
			numOfWinners++
		}
	}

	return numOfWinners
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

func getWinners(boards []board) []board {
	// marked values should fill a column or row in the board
	winners := make([]board, 0)
	for i := range boards {

		if boards[i].isWinner {
			continue
		}

		var markedColumns [5]int
		for y := range boards[i].value {
			var markedRow int
			for x := range boards[i].value[y] {

				if !boards[i].value[y][x].isMarked {
					continue
				}

				markedRow++
				markedColumns[x]++

				// if any marked rows or columns have 5 then thats a winner
				if markedRow >= 5 || markedColumns[x] >= 5 {
					boards[i].isWinner = true
					// this is the last winner

					winners = append(winners, boards[i])
				}
			}
		}
	}

	return winners
}

func markBoard(board [][]boardTile, value string) {
	for k := range board {
		for i := range board[k] {
			if board[k][i].isMarked || board[k][i].value != value {
				continue
			}

			board[k][i].isMarked = true
			return
		}
	}
}

func buildBoards(boardInputs []string, numOfBoards int) []board {

	boards := make([]board, numOfBoards)

	for i, boardInput := range boardInputs {
		boardInputLines := strings.Split(boardInput, "\n")

		boards[i] = board{
			value:    make([][]boardTile, 5),
			isWinner: false,
		}

		for y, boardInputLine := range boardInputLines {

			boardInputTiles := strings.Split(boardInputLine, " ")

			for _, boardInputTile := range boardInputTiles {

				// one digit values have an space
				// in front of them so we must get rid of empty values
				if boardInputTile == " " || boardInputTile == "" {
					continue
				}

				boards[i].value[y] = append(
					boards[i].value[y],
					boardTile{
						value:    boardInputTile,
						isMarked: false,
					})
			}
		}
	}

	return boards
}
