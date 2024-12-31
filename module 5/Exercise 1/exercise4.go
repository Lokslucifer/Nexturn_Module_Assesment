package main

import (
	"fmt"
	"strconv"
)

type Question struct {
	QuestionText string
	Options      [4]string
	Answer       int
}

var questionBank = []Question{
	{
		QuestionText: "What is the capital of India?",
		Options:      [4]string{"1. Berlin", "2. New Delhi", "3. Paris", "4. Rome"},
		Answer:       3,
	},
	{
		QuestionText: "Which planet is known as the Red Planet?",
		Options:      [4]string{"1. Earth", "2. Mars", "3. Jupiter", "4. Venus"},
		Answer:       2,
	},
	{
		QuestionText: "What is the largest ocean on Earth?",
		Options:      [4]string{"1. Atlantic", "2. Indian", "3. Arctic", "4. Pacific"},
		Answer:       4,
	},
}

func main() {
	var score int
	var answerInput string

	for i, question := range questionBank {
		fmt.Printf("Question %d: %s\n", i+1, question.QuestionText)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		fmt.Print("Enter your answer (1-4) or type 'exit' to quit: ")
		fmt.Scanln(&answerInput)

		if answerInput == "exit" {
			fmt.Println("Exiting the quiz...")
			break
		}

		answer, err := strconv.Atoi(answerInput)
		if err != nil || answer < 1 || answer > 4 {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		if answer == question.Answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Wrong!")
		}
	}

	fmt.Printf("\nYour score: %d/%d\n", score, len(questionBank))

	if score == len(questionBank) {
		fmt.Println("Excellent!")
	} else if score >= len(questionBank)/2 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Needs Improvement.")
	}
}
