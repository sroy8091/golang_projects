package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "Please specify a problem set in format of (question, answer)")
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

	for i, p := range problem {
		var answer string
		fmt.Printf("Problem #%d :  %s\n", i+1, p.q)

		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
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
