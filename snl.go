package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var l90 = []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91}
var l80 = []int{81, 82, 83, 84, 85, 86, 87, 88, 89, 90}
var l70 = []int{80, 79, 78, 77, 76, 75, 74, 73, 72, 71}
var l60 = []int{61, 62, 63, 64, 65, 66, 67, 68, 69, 70}
var l50 = []int{60, 59, 58, 57, 56, 55, 54, 53, 52, 51}
var l40 = []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
var l30 = []int{40, 39, 38, 37, 36, 35, 34, 33, 32, 31}
var l20 = []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var l10 = []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11}
var l00 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var numlist = [][]int{l90, l80, l70, l60, l50, l40, l30, l20, l10, l00}
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
		for userScore < 100 && compScore < 100 {
			userDice := rollDice()
			userScore += userDice
			if valInSlice(userScore) == 1 {
				fmt.Printf("User hit a snake! Score: %d -> %d.\n", userScore, snakes[userScore])
				userScore = snakes[userScore]
			} else if valInSlice(userScore) == 2 {
				fmt.Printf("User hit a ladder! Score: %d -> %d.\n", userScore, ladders[userScore])
				userScore = ladders[userScore]
			}
			compDice := rollDice()
			compScore += compDice
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
