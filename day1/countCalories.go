package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	overallHighest := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	currentCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Println("New line, reset count")
			if currentCount > overallHighest {
				fmt.Printf("current is: %v \n", currentCount)
				fmt.Printf("highest is: %v \n", overallHighest)
				overallHighest = currentCount
			}
			currentCount = 0
		} else {
			asInt, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(asInt)
			currentCount = currentCount + asInt
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
