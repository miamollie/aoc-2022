package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		firstRange := strings.Split(sections[0], "-")
		secondRange := strings.Split(sections[1], "-")
		if containsPartialRange(firstRange, secondRange) || containsPartialRange(secondRange, firstRange) {
			fmt.Printf("firstRange %v, secondRange %v\n", firstRange, secondRange)
			total++
		}

	}
	fmt.Printf("total cost %v\n", total)
}

func containsFullRange(a, b []string) bool {
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])
	return a0 >= b0 && a1 <= b1
}

func containsPartialRange(a, b []string) bool {
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])

	arr := makeArray(a0, a1)

	for _, v := range arr {
		if v >= b0 && v <= b1 {
			return true
		}
	}
	return false
}

func makeArray(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}
