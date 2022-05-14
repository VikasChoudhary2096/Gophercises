package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type quesAns struct {
	ques string
	ans  string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "question answer csv file")
	timeLimit := flag.Int("t", 30, "time limit of quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(file)
	content, _ := csvReader.ReadAll()
	// var key string
	fmt.Println("Give answers to the questions...")
	// fmt.Scan(&key)

	var result int

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

pairLoop:
	for _, pair := range content {
		fmt.Printf("what is %v?", pair[0])
		fmt.Println()
		answer := make(chan string)

		go func() {
			var ans string
			fmt.Scan(&ans)
			answer <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println("time is over")
			break pairLoop
			// pairLoop.break
		case ans := <-answer:
			if ans == strings.TrimSpace(pair[1]) {
				result += 1
			}
		}
	}
	fmt.Println()
	fmt.Printf("your correct answers count is %d", result)
	fmt.Println()
}
