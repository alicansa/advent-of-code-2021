package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const SimulationLength = 256

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

	// store the number of fish in given internal timer state
	// this way we don't need to append the list
	var lanternfishInternalTimerStates [9]int

	// initialise
	for _, initialState := range initialStateInput {
		state, err := strconv.Atoi(initialState)
		if err != nil {
			log.Fatal(err)
			return
		}

		lanternfishInternalTimerStates[state]++
	}

	for dayCounter := 1; dayCounter <= SimulationLength; dayCounter++ {
		lanternfishInternalTimerStates = processInternalTimerStates(lanternfishInternalTimerStates)
	}

	var totalNumOfFish int

	for _, state := range lanternfishInternalTimerStates {
		totalNumOfFish += state
	}

	println(totalNumOfFish)
}

func processInternalTimerStates(lanternfishInternalTimerStates [9]int) [9]int {

	var newLanternFishInternalTimerState [9]int

	for i := 8; i >= 0; i-- {
		if lanternfishInternalTimerStates[i] <= 0 {
			continue
		}

		switch i {
		// gives birth
		// remove fish from 0
		// add to 6 and set to 8
		case 0:
			newLanternFishInternalTimerState[6] = newLanternFishInternalTimerState[6] + lanternfishInternalTimerStates[0]
			newLanternFishInternalTimerState[8] = lanternfishInternalTimerStates[0]
			newLanternFishInternalTimerState[0] = lanternfishInternalTimerStates[1]
			break
		default:
			newLanternFishInternalTimerState[i-1] = lanternfishInternalTimerStates[i]
			break
		}
	}

	return newLanternFishInternalTimerState
}
