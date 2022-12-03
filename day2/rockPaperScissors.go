package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Result string
type Move string

/*
X: 6  -> Win
Y: 3  -> Lose
Z: 0  -> Draw
*/

const (
	ResultWin  Result = "Z"
	ResultLose Result = "X"
	ResultDraw Result = "Y"
)

var outcomeScores = map[Result]int{ResultWin: 6, ResultDraw: 3, ResultLose: 0}

/*
A: 1  -> Rock
B: 2  -> Paper
C: 3  -> Scissors
*/

const (
	MoveRock     Move = "A"
	MovePaper    Move = "B"
	MoveScissors Move = "C"
)

var moveScores = map[Move]int{MoveRock: 1, MovePaper: 2, MoveScissors: 3}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		hands := strings.Split(line, " ")

		ownOutcome := Result(hands[1])
		opponentHand := Move(hands[0])

		totalScore += outcomeScores[ownOutcome]
		ownHand := roundHandFromOutcome(ownOutcome, opponentHand)
		totalScore += moveScores[ownHand]
	}
	fmt.Println(totalScore)
}

// func roundOutcome(self Move, opponent Move) Result {
// 	switch self {
// 	case "X":
// 		if opponent == "A" {
// 			return ResultDraw
// 		} else if opponent == "B" {
// 			return ResultLose
// 		} else {
// 			return ResultWin
// 		}
// 	case "Y":
// 		if opponent == "B" {
// 			return ResultDraw
// 		} else if opponent == "C" {
// 			return ResultLose
// 		} else {
// 			return ResultWin
// 		}

// 	case "Z":
// 		if opponent == "C" {
// 			return ResultDraw
// 		} else if opponent == "A" {
// 			return ResultLose
// 		} else {
// 			return ResultWin
// 		}
// 	default:
// 		panic("uh oh")
// 	}
// }

//	y = draw, z = win, x = lose
//
// `A for Rock, B for Paper, and C for Scissors`
func roundHandFromOutcome(desiredOutcome Result, opponentHand Move) Move {
	switch desiredOutcome {
	case ResultDraw:
		return opponentHand
	case ResultWin:
		if opponentHand == MoveRock {
			return MovePaper
		} else if opponentHand == MovePaper {
			return MoveScissors
		} else {
			return MoveRock
		}
	case ResultLose:
		if opponentHand == MoveRock {
			return MoveScissors
		} else if opponentHand == MovePaper {
			return MoveRock
		} else {
			return MovePaper
		}
	default:
		panic("uh oh")
	}
}
