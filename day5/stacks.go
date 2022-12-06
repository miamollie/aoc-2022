package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stacksFilePath := os.Args[1]
	stacksFile, err := os.Open(stacksFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer stacksFile.Close()

	scanner := bufio.NewScanner(stacksFile)
	stacks := make([]string, 9)
	for scanner.Scan() {
		line := scanner.Text()
		addToStacks(line, stacks)
	}

	fmt.Printf("STACKS FIRST POSITION")
	fmt.Printf("%#v\n", stacks)

	movesFilePath := os.Args[2]
	movesFile, err := os.Open(movesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer movesFile.Close()

	scanner = bufio.NewScanner(movesFile)
	for scanner.Scan() {
		line := scanner.Text()
		applyMultiBoxMove(line, stacks)
	}
	fmt.Printf("STACKS FINAL POSITION")
	fmt.Printf("%#v\n", stacks)
}

func applySingleBoxMove(line string, stacks []string) {
	sections := strings.Split(line, " ")
	boxCount, _ := strconv.Atoi(sections[1])
	fromColNum, _ := strconv.Atoi(sections[3])
	fromColNum-- //array offset
	toColNum, _ := strconv.Atoi(sections[5])
	toColNum-- //array offset

	fmt.Printf("move  %v  boxes from %v,  to %v\n", boxCount, fromColNum+1, toColNum+1)
	for i := 0; i < boxCount; i++ {
		fmt.Printf("%#v\n", stacks)
		oldCol := stacks[fromColNum]
		newCol := stacks[toColNum]
		stacks[toColNum] = oldCol[0:1] + newCol
		l := len(oldCol)
		stacks[fromColNum] = oldCol[1:l]
	}

}

func applyMultiBoxMove(line string, stacks []string) {
	sections := strings.Split(line, " ")
	boxCount, _ := strconv.Atoi(sections[1])
	fromColNum, _ := strconv.Atoi(sections[3])
	fromColNum-- //array offset
	toColNum, _ := strconv.Atoi(sections[5])
	toColNum-- //array offset

	fmt.Printf("move  %v  boxes from %v,  to %v\n", boxCount, fromColNum+1, toColNum+1)

	oldCol := stacks[fromColNum]
	newCol := stacks[toColNum]
	stacks[toColNum] = oldCol[0:boxCount] + newCol
	l := len(oldCol)
	stacks[fromColNum] = oldCol[boxCount:l]

}

func addToStacks(line string, stacks []string) {
	t := strings.ReplaceAll(line, "    ", " ***")
	fmt.Println(t)
	re := regexp.MustCompile(`\[(.)\]|(\*\*\*)`)
	x := re.FindAllStringSubmatch(t, -1)

	for i, v := range x {
		// fmt.Printf("stack %#v with value %#v \n", i, v[1])
		newBlock := v[1]
		curStack := stacks[i]
		if newBlock != "" {
			// fmt.Printf("stack number %v contents: %#v \n", i+1, v)
			stacks[i] = curStack + newBlock
		}
	}
}
