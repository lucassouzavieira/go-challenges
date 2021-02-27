package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	Question string
	Answer   string
}

func (q Question) CheckAnswer(aswer string) bool {
	return q.Answer == aswer
}

func main() {
	fmt.Println("Quiz Game!")
	fmt.Println("Answer the following questions")

	// Open CSV file
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println(err)
	}

	// Close CSV file at the end of the program
	defer csvFile.Close()

	reader, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println("Failed to load CSV file")
	}

	// Correct answers
	var correct int = 0

	for index, line := range reader {
		// Skip CSV headers
		if index == 0 {
			continue
		}

		question := Question{
			Question: line[0],
			Answer:   line[1],
		}

		fmt.Println(question.Question)

		var answer string
		fmt.Scan(&answer)

		if question.CheckAnswer(answer) {
			correct = correct + 1
		}

		fmt.Println("")
	}

	fmt.Println("You answered " + fmt.Sprint(correct) + " questions correctly!")
}
