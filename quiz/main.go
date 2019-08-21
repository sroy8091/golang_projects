package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "Please specify a problem set in format of (question, answer)")
	limit := flag.Int64("limit", 15, "Please specify a time limit for the quiz")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Falied to parse the file")
		os.Exit(1)
	}

	problem := parseLine(lines)

	correct := 0

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, p := range problem {
		answerCh := make(chan string)
		fmt.Printf("Problem #%d :  %s\n", i+1, p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problem))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problem))

}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
