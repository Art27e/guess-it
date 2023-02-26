package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var seconds = time.Now().Unix()
var target int

func menu() int {
	fmt.Println("Who is going to guess [PC/User]")
	readChoice := bufio.NewReader(os.Stdin)
	choiceInput, err := readChoice.ReadString('\n') // read data before pressing ENTER
	if err != nil {
		log.Fatal(err)
	}
	choiceInput = strings.TrimSpace(choiceInput) // removing newline
	if choiceInput == "1" {
		return 1
	}
	if choiceInput == "2" {
		return 2
	}
	return 0
}

func main() {
	checkChoice := menu()
	if checkChoice == 1 {
		pcGuess()
	}
	if checkChoice == 2 {
		userGuess()
	}
	if checkChoice == 0 {
		os.Exit(2)
	}
}

func userGuess() {
	rand.Seed(seconds)
	target = rand.Intn(100-1) + 1 // range 1-100
	fmt.Println("I have choosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false // default
myLoop:
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case guess > target:
			fmt.Println("Your guess was HIGH")
		case guess < target:
			fmt.Println("Your guess was LOW")
		default:
			success = true // if user guesses right, no message of fault guess
			fmt.Println("Exactly! You are right!")
			break myLoop
		}
	}
	if !success {
		fmt.Println("Sorry. You didn't guess my number. It was:", target)
	}
}

func pcGuess() {
	rand.Seed(seconds)
	target = rand.Intn(100-1) + 1 // range 1-100
	fmt.Println("PC, can you guess it? Range [1-100]")
	myNumber := bufio.NewReader(os.Stdin)
	myNumberForPC, err := myNumber.ReadString('\n') // read data before pressing ENTER
	if err != nil {
		log.Fatal(err)
	}
	myNumberForPC = strings.TrimSpace(myNumberForPC)
	myNumberForPCint, err := strconv.Atoi(myNumberForPC)
	success := false
	highSlice := []int{}
	lowSlice := []int{}
myLoop:
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("PC, You have", 10-guesses, "guesses left")
		fmt.Print("Make a guess: ")
		guess := target
		fmt.Println(guess)
		switch {
		case guess > myNumberForPCint:
			highSlice = append(highSlice,guess)
			if len(lowSlice) == 0 {
				lowSlice = append(lowSlice,1)
			}
			target = rand.Intn(highSlice[len(highSlice)-1]-2-lowSlice[len(lowSlice)-1]+1) + lowSlice[len(lowSlice)-1]+1 // 64-1 57-1 // replaced guess-1 (max)
			fmt.Println("Your guess was HIGH")
		case guess < myNumberForPCint:
			lowSlice = append(lowSlice,guess)
			if len(highSlice) == 0 {
				highSlice = append(highSlice,100)
			}
			target = rand.Intn(highSlice[len(highSlice)-1]-1-lowSlice[len(lowSlice)-1]+1) + lowSlice[len(lowSlice)-1]+1 // rand.Intn(MAX-MIN) + MIN // replaced guess+1 (MIN)
			fmt.Println("Your guess was LOW")
		default:
			success = true // if user guesses right, no message of fault guess
			fmt.Println("Exactly! You are right!")
			break myLoop
		}
	}
	if !success {
		fmt.Println("Sorry, PC. You didn't guess my number. It was:", myNumberForPCint)
	}
}