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
	totalLength := len(bitGroups)
	lengthPerGroup := len(bitGroups[0])
	numberOfOnes := make([]int, lengthPerGroup)

	for _, bits := range bitGroups {
		for k, bit := range bits {
			if bit == '1' {
				numberOfOnes[k]++
			}
		}
	}

	var mcb string
	var lcb string
	for _, item := range numberOfOnes {
		if item > totalLength/2 {
			mcb = mcb + "1"
			lcb = lcb + "0"
			continue
		}

		mcb = mcb + "0"
		lcb = lcb + "1"
	}

	//convert string to number
	gammaRate, err := convertToBase10(mcb)
	epsilonRate, err := convertToBase10(lcb)

	if err != nil {
		log.Fatal(err)
		return
	}

	println(gammaRate * epsilonRate)
}

func convertToBase10(bits string) (int64, error) {
	return strconv.ParseInt(bits, 2, 0)
}
