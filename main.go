package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	//os.Open opens a file, asterisk = pointer to a string
	file, err := os.Open(*csvFile)
	//if there's an error when opening the file log the output through the err keyword
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", err)
		//if anything other than 0, then exit
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
