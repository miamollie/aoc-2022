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
	errorCost := 0
	for scanner.Scan() {
		line := scanner.Text()
		compartmentLength := len(line) / 2
		var foundItems = make(map[rune]struct{})

		for i, v := range line {
			fmt.Printf("Letter is currently %v \n", string(v))
			if i < compartmentLength {
				fmt.Println("still on left side")
				foundItems[v] = struct{}{}
			} else {
				fmt.Println("still on now on right side")
				if _, ok := foundItems[v]; ok {
					fmt.Printf("%v has been seen before and index is %v \n", string(v), i)
					fmt.Printf(" Priority cost for letters is %v\n", strings.Index(letters, string(v)))
					errorCost += strings.Index(letters, string(v)) + 1
					fmt.Printf("new total %v,  letter Cost %v, letter %v \n", errorCost, strings.Index(string(v), letters)+1, string(v))
					break
				}
			}
		}
	}
	fmt.Println(errorCost)
}
