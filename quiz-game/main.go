package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	var filename string
	var timeLimit int
	flag.StringVar(&filename, "csv", "problems.csv", "Csv file for quiz.")
	flag.IntVar(&timeLimit, "time", 30, "Time Limit for quiz game.")
	flag.Parse()

	csvData := ReadCsv(filename)
	questionList := CreateQuestionList(csvData)
	PlayTheGame(questionList)

}

func ReadCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data

}

func CreateQuestionList(data [][]string) []Quiz {
	var quizzes []Quiz

	for _, line := range data {
		q := Quiz{question: strings.TrimSpace(line[0]), answer: strings.TrimSpace(line[1])}

		quizzes = append(quizzes, q)
	}

	return quizzes
}

func PlayTheGame(questions []Quiz) {
	point := 0
	reader := bufio.NewReader(os.Stdin)
	for _, question := range questions {
		fmt.Printf("%s = ", question.question)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)
		if answer == question.answer {
			point += 1
		}
	}

	fmt.Printf("Point: %d\n", point)

}
