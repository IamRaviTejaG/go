package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var snakes = map[int]int{29: 7, 38: 20, 44: 14, 55: 11, 73: 52, 82: 60, 93: 43, 96: 17, 98: 48}
var ladders = map[int]int{3: 21, 4: 36, 15: 48, 24: 58, 31: 70, 49: 90, 60: 79, 63: 99, 72: 91, 77: 97}
var snakeKeys = make([]int, 0, len(snakes))
var ladderKeys = make([]int, 0, len(ladders))

func rollDice() int {
	diceRoll := rand.Intn(6) + 1
	return diceRoll
}

func slKeys() {
	for k := range snakes {
		snakeKeys = append(snakeKeys, k)
	}
	for k := range ladders {
		ladderKeys = append(ladderKeys, k)
	}
}

func valInSlice(a int) int {
	for _, b := range snakeKeys {
		if b == a {
			return 1
		}
	}
	for _, b := range ladderKeys {
		if b == a {
			return 2
		}
	}
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	slKeys()                     // Generate a slice of snake keys & ladder keys
	userScore, compScore := 0, 0 // Initial scores
	fmt.Print("Enter y to roll dice: ")
	var text string
	fmt.Scanln(&text)
	for text == "y" || text == "Y" {
		for userScore < 101 || compScore < 101 {
			userDice := rollDice()
			if userScore+userDice > 100 {
				userScore += 0
			} else {
				userScore += userDice
			}
			if valInSlice(userScore) == 1 {
				fmt.Printf("User hit a snake! Score: %d -> %d.\n", userScore, snakes[userScore])
				userScore = snakes[userScore]
			} else if valInSlice(userScore) == 2 {
				fmt.Printf("User hit a ladder! Score: %d -> %d.\n", userScore, ladders[userScore])
				userScore = ladders[userScore]
			}
			compDice := rollDice()
			if compScore+compDice > 100 {
				compScore += 0
			} else {
				compScore += compDice
			}
			if valInSlice(compScore) == 1 {
				fmt.Printf("Computer hit a snake! Score: %d -> %d.\n", compScore, snakes[compScore])
				compScore = snakes[compScore]
			} else if valInSlice(compScore) == 2 {
				fmt.Printf("Computer hit a ladder! Score: %d -> %d.\n", compScore, ladders[compScore])
				compScore = ladders[compScore]
			}
			fmt.Printf("userDice: %d\n", userDice)
			fmt.Printf("compDice: %d\n", compDice)
			fmt.Printf("userScore: %d\n", userScore)
			fmt.Printf("compScore: %d\n", compScore)
			if userScore >= 100 {
				fmt.Println("USER WINS!")
				os.Exit(0)
			} else if compScore >= 100 {
				fmt.Println("COMPUTER WINS!")
				os.Exit(0)
			} else {
				fmt.Print("Enter y to roll dice: ")
				fmt.Scanln(&text)
			}
		}
	}
}
