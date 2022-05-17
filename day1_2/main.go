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

	ss := strings.Split(string(f), "\n")
	previousAverage := -1
	currentAverage := 0
	increaseCounter := 0

	for i := range ss {

		if i == len(ss)-2 {
			break
		}

		currentAverage, err = calculateWindowAverage(ss[i], ss[i+1], ss[i+2])

		if err != nil {
			log.Fatal(err)
			break
		}

		if previousAverage >= 0 && currentAverage > previousAverage {
			increaseCounter++
		}

		previousAverage = currentAverage
	}

	println(increaseCounter)
}

func calculateWindowAverage(s1 string, s2 string, s3 string) (int, error) {
	i1, err := strconv.Atoi(s1)
	i2, err := strconv.Atoi(s2)
	i3, err := strconv.Atoi(s3)

	if err != nil {
		return -1, err
	}

	return (i1 + i2 + i3), err
}
