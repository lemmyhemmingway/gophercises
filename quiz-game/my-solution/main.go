package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

const (
	DEFAULT_TIME     = 30
	DEFAULT_CSV_FILE = "problems.csv"
)

var point int

func main() {
	var filename string
	var timeLimit int
	flag.StringVar(&filename, "file", DEFAULT_CSV_FILE, "Csv file for quiz.")
	flag.IntVar(&timeLimit, "limit", DEFAULT_TIME, "Time Limit for quiz game.")
	flag.Parse()

	csvData := ReadCsv(filename)
	questionList := CreateQuestionList(csvData)
	timer1 := time.NewTimer(time.Duration(timeLimit) * time.Second)
	go func() {
		PlayTheGame(questionList)
	}()
	<-timer1.C
	fmt.Println()
	fmt.Printf("%d\n", point)

	// ? TODO: stop channel when question list done

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
	// var quizzes []Quiz // already knew the size do not use append
	quizzes := make([]Quiz, len(data))

	for i, line := range data {
		quizzes[i] = Quiz{question: strings.TrimSpace(line[0]), answer: strings.TrimSpace(line[1])}

		// quizzes = append(quizzes, q)
	}

	return quizzes
}

func PlayTheGame(questions []Quiz) {
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
