package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var increaseCounter int
	// open data file
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// defer closing until the end of this function
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var currentDepth int
	previousDepth := -1

	for scanner.Scan() {

		//read each line
		var stringDepth = scanner.Text()

		currentDepth, err = strconv.Atoi(stringDepth)
		if err != nil {
			log.Fatal(err)
			break
		}

		if currentDepth > previousDepth && previousDepth >= 0 {
			increaseCounter++
		}

		previousDepth = currentDepth
	}

	//log if there is an error reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increaseCounter)
}
