package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//function to check if there are completely empty spaces in csv file -> returns an error if either the first or second index on a line is empty
func sliceFormat(ls [][]string) ([][]string, error) {
	for _, item := range ls {
		if strings.TrimSpace(item[0]) == "" || strings.TrimSpace(item[1]) == "" {
			return ls, errors.New("Incorrect database format.")
		}
	}
	return ls, nil
}

func main() {

	//flags:
	//fname -> name of csv file, default is problems.csv
	//probCount -> number of questions to ask, default is 10
	fname := flag.String("csv", "problems.csv", "problems database")
	probCount := flag.Int("n", 10, "integer")
	flag.Parse()

	//read csv, flag errors
	file, err := os.Open(*fname)
	if err != nil {
		log.Fatalf("failed to open csv file: %s", err)
	} else {
		fmt.Println("Successfully opened csv file!")
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed to parse csv file: %s", err)
	} else {
		fmt.Println("Successfully parsed csv file!")
	}
	lines, err1 := sliceFormat(lines)
	if err1 != nil {
		log.Fatalf("Incorrect file format: %s", err1)
	}

	//randomly shuffle list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })

	//remove extra questions from list if probCount is less than the length of lines in the csv
	//if there are not enough questions available, only ask available questions
	if *probCount <= len(lines) {
		lines = lines[:*probCount]
	} else {
		fmt.Printf("Insufficient number of questions. Only %v questions will be asked.", len(lines))
		*probCount = len(lines)
	}

	//iterate and scan
	var correctCount int
	for idx, item := range lines {
		var ans string
		fmt.Printf("Q%v: %v = ", idx+1, strings.TrimSpace(item[0]))
		fmt.Scan(&ans)
		if ans == strings.TrimSpace(item[1]) {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Println("Incorrect!")
		}
	}
	fmt.Printf("You answered %v out of %v questions correctly.\n", correctCount, *probCount)
}
