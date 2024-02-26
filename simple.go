package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my riddle game!")
	fmt.Printf("What's your name? ")
	var name string
	var age int
	var answer1 string
	var answer2 string
	var answer3 string

	fmt.Scan(&name)

	fmt.Printf("Welcome to the game...%v, I hope you are ready for what is to come\n", name)

	fmt.Printf("Before we proceed, what's your age? ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Good you seem old enough to play ")
	} else {
		fmt.Println("Hmm you don't seem old enough to be toyed with, I'll come back when you are...more ripe ")
		return
	}

	fmt.Println("This quiz games has three riddles in it, if you get all three you pass if you don't you fail(die)")

	fmt.Printf("Here is the first riddle. What is at the end of a rainbow? ")
	fmt.Scan(&answer1)

	if strings.EqualFold(answer1, "W") {
		fmt.Println("Hmm...you got it, you safe for now")
	} else {
		fmt.Println("You are wrong, therefore dead")
		return
	}

	fmt.Printf("Here is the second riddle. What kind of goose fights with snakes? ")
	fmt.Scan(&answer2)

	if strings.EqualFold(answer2, "mongoose") {
		fmt.Println("Hmm...you got it, that is three. I don't like this. You safe for now")
	} else {
		fmt.Println("You are wrong, therefore dead")
		return
	}

	fmt.Printf("Here is the third riddle. What has an eye but can not see? ")
	fmt.Scan(&answer3)

	if strings.EqualFold(answer3, "needle") {
		fmt.Println("This leaves a bad taste in my mouth, but you win")
	} else {
		fmt.Println("Oh so you thought you would win, fucking loser...now time for you to die")
		return
	}

}
