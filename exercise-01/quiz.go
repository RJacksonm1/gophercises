package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	questionsFile := flag.String("csv", "problems.csv", "The questions for the quiz. One question per row. Question first, answer second.")
	flag.Parse()

	var total, right int16

	f, err := os.Open(*questionsFile)
	check(err)

	userInputReader := bufio.NewReader(os.Stdin)

	r := csv.NewReader(bufio.NewReader(f))
	for {
		qa, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		q, a := qa[0], strings.ToLower(qa[1])

		fmt.Printf("Question %d: %s\n", total+1, q)
		total++

		userAnswer, _ := userInputReader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)
		userAnswer = strings.ToLower(userAnswer)

		if userAnswer == a {
			right++
		}
	}

	fmt.Printf("You got %d out of %d questions correct!\n", right, total)
}
