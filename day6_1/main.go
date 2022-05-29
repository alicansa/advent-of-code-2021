package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const SimulationLength = 80

func main() {
	//get inputs
	f, err := os.ReadFile("data.txt")

	// log error
	if err != nil {
		log.Fatal(err)
		return
	}

	// get the initial state
	initialStatesInput := strings.Split(string(f), "\n")[0]

	initialStateInput := strings.Split(initialStatesInput, ",")
	lanternfishInternalTimerStates := make([]int, len(initialStateInput))

	// initialise
	for i, initialState := range initialStateInput {
		state, err := strconv.Atoi(initialState)
		if err != nil {
			log.Fatal(err)
			return
		}

		lanternfishInternalTimerStates[i] = state
	}

	for dayCounter := 1; dayCounter <= SimulationLength; dayCounter++ {
		lanternfishInternalTimerStates = processInternalTimerStates(lanternfishInternalTimerStates)
	}

	println(len(lanternfishInternalTimerStates))
}

func processInternalTimerStates(lanternfishInternalTimerStates []int) []int {

	for i := range lanternfishInternalTimerStates {
		if lanternfishInternalTimerStates[i] == 0 {
			lanternfishInternalTimerStates[i] = 6
			lanternfishInternalTimerStates = append(lanternfishInternalTimerStates, 8)
			continue
		}

		lanternfishInternalTimerStates[i]--
	}

	return lanternfishInternalTimerStates
}
