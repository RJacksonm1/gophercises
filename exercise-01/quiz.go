package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Question to ask in the quiz
type Question struct {
	Q string
	A string
}

// CheckAnswer verifies the user's answer to this question
func (q *Question) CheckAnswer(a string) bool {
	return strings.TrimSpace(strings.ToLower(a)) == strings.ToLower(q.A)
}

// Questions to ask in the quiz
type Questions []Question

// LoadQuestionsFromFile creates a new Questions from a CSV of question and answer pairs
func LoadQuestionsFromFile(questionsFile string) (q Questions, err error) {
	f, err := os.Open(questionsFile)
	if err != nil {
		return Questions{}, err
	}

	r := csv.NewReader(bufio.NewReader(f))
	for {
		qa, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Questions{}, err
		}

		q = append(q, Question{
			Q: qa[0],
			A: strings.TrimSpace(qa[1]),
		})
	}
	return q, err
}

// Shuffle the order of the questions
func (qs *Questions) Shuffle() {
	rand.Shuffle(len(*qs), func(i, j int) {
		(*qs)[i], (*qs)[j] = (*qs)[j], (*qs)[i]
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	questionsFile := flag.String("csv", "problems.csv", "The questions for the quiz. One question per row. Question first, answer second.")
	shuffle := flag.Bool("shuffle", false, "Randomise the order of the questions")
	limit := flag.Int("limit", 30, "Set a time limit on the quiz")
	flag.Parse()

	questions, err := LoadQuestionsFromFile(*questionsFile)
	check(err)

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		questions.Shuffle()
	}

	finished := make(chan bool)
// 
	// Ask questions
	var total, right int16
	go func() {
		userInputReader := bufio.NewReader(os.Stdin)
		for _, q := range questions {
			fmt.Printf("Question %d: %s\n", total+1, q.Q)
			total++

			a, _ := userInputReader.ReadString('\n')
			if q.CheckAnswer(a) {
				right++
			}
		}
// 
		finished <- true
	}()

	// Time limitt
	go func() {
		time.Sleep(time.Duration(*limit) * time.Second)
		fmt.Println("\nTime's up!")
		finished <- true
	}()

	if <-finished {
		fmt.Printf("You got %d out of %d questions correct!\n", right, total)
	}
}
