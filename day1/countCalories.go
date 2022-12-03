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

	first, second, third := 0, 0, 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	currentCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Println("New line, reset count")
			if currentCount > first {
				third = second
				second = first
				first = currentCount
			} else if currentCount > second {
				third = second
				second = currentCount
			} else if currentCount > third {
				third = currentCount
			}

			fmt.Printf("current is: %v \n", currentCount)
			fmt.Printf("highest is: %v %v %v \n", first, second, third)
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

	total := first + second + third
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
