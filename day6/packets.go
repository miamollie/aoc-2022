package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const offset = 14

func main() {
	stacksFilePath := os.Args[1]
	stacksFile, err := os.Open(stacksFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer stacksFile.Close()

	scanner := bufio.NewScanner(stacksFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line %#v is %#v chars long\n", line, len(line))

		for i := 0; i < len(line)-offset; i = i + 1 {
			// check for a buffer of 4 unique characters
			// move through the input string in slices of 4
			cur := line[i : i+offset]
			fmt.Printf("Checking range %#v \n", cur)
			if ok := isUniqueRange(cur); ok {
				fmt.Println(i + offset)
				break
			}

		}
		fmt.Println("Finished looking!")
	}
}

func isUniqueRange(s string) bool {
	chars := make(map[string]struct{}, offset)
	for j := 0; j < offset; j++ {
		letter := string(s[j])
		// fmt.Printf("Letter is %#v \n", letter)
		if _, ok := chars[letter]; ok {
			// fmt.Printf("Letter has been seen before; exit to check new range\n")
			return false
		} else {
			chars[letter] = struct{}{}
			// fmt.Printf("added letter to chars; %#v \n", chars)
		}
	}
	return true // if we get through 4 chars without seeing a duplicate, they're all unique!
}
