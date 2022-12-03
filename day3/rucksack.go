package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCost := 0
	groupCount := 0
	// itemsMap := make(map[string][]int)
	for scanner.Scan() {
		groupCount++
		line := scanner.Text()
		updatePriorityCost(line, &totalCost)
		// for _, v := range line {
		// 	c := string(v)
		// 	val, ok := itemsMap[c]

		// 	fmt.Printf("value for letter  %v is %v and groupCount is %v\n ", c, itemsMap[c], groupCount)
		// 	if !ok {
		// 		fmt.Println(val)
		// 		arr := make([]int, 3)
		// 		arr[groupCount-1] = 1
		// 		itemsMap[c] = arr
		// 	} else {
		// 		val[groupCount-1] = 1
		// 		itemsMap[c] = val
		// 	}
		// 	// found in all three groups
		// 	if len(val) == 3 && val[0]+val[1]+val[2] == 3 {
		// 		fmt.Printf("Found item %v 3 times and it costs %v \n", c, strings.Index(letters, c))
		// 		totalCost += strings.Index(letters, c) + 1
		// 		groupCount = 0
		// 		itemsMap = make(map[string][]int)
		// 		break
		// 	}
		// }
	}
	fmt.Printf("total cost %v\n", totalCost)
}

func updatePriorityCost(line string, errorCost *int) {
	compartmentLength := len(line) / 2
	var foundItems = make(map[rune]struct{})

	for i, v := range line {
		if i < compartmentLength {
			foundItems[v] = struct{}{}
		} else {
			if _, ok := foundItems[v]; ok {
				*errorCost += strings.Index(letters, string(v)) + 1
				break
			}
		}
	}
}
