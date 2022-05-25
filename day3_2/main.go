package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")

	// log error
	if err != nil {
		log.Fatal(err)
		return
	}

	bitGroups := strings.Split(string(f), "\n")

	mcb := filterBitsStartingWithCondition(
		bitGroups,
		0,
		func(filteredBitGroups []string, bits string, bitPosition int) bool {
			mcbIsOne := isMCBOne(filteredBitGroups, bitPosition)
			return (mcbIsOne && bits[bitPosition] == '1') || (!mcbIsOne && bits[bitPosition] == '0')
		})

	lcb := filterBitsStartingWithCondition(
		bitGroups,
		0,
		func(filteredBitGroups []string, bits string, bitPosition int) bool {
			lcbIsOne := isLCBOne(filteredBitGroups, bitPosition)
			return (lcbIsOne && bits[bitPosition] == '1') || (!lcbIsOne && bits[bitPosition] == '0')
		})

	//convert string to number
	oxygenGeneratorRating, err := convertToBase10(mcb)
	co2ScrubberRating, err := convertToBase10(lcb)

	if err != nil {
		log.Fatal(err)
		return
	}

	println(oxygenGeneratorRating * co2ScrubberRating)
}

func filterBitsStartingWithCondition(
	bitGroups []string,
	bitPosition int,
	condition func(bitGroups []string, bits string, bitPosition int) bool) string {
	if len(bitGroups) == 1 {
		return bitGroups[0]
	}

	var filteredGroups []string

	for _, bits := range bitGroups {
		if !condition(bitGroups, bits, bitPosition) {
			continue
		}

		filteredGroups = append(filteredGroups, bits)
	}

	return filterBitsStartingWithCondition(filteredGroups, bitPosition+1, condition)
}

func isMCBOne(bitGroups []string, bitPosition int) bool {
	totalLength := len(bitGroups)
	numberOfOnes := getNumberOfOnes(bitGroups, bitPosition)

	return 2*numberOfOnes >= totalLength
}

func isLCBOne(bitGroups []string, bitPosition int) bool {
	totalLength := len(bitGroups)
	numberOfOnes := getNumberOfOnes(bitGroups, bitPosition)

	return 2*numberOfOnes < totalLength
}

func getNumberOfOnes(bitGroups []string, bitPosition int) int {
	var numberOfOnes int

	for _, bits := range bitGroups {
		if bits[bitPosition] == '1' {
			numberOfOnes++
		}
	}

	return numberOfOnes
}

func convertToBase10(bits string) (int64, error) {
	return strconv.ParseInt(bits, 2, 0)
}
