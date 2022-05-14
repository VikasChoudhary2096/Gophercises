package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type quesAns struct {
	ques string
	ans  int
}

func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}

	content := make([]byte, 75)

	_, e := file.Read(content)
	if err != nil {
		fmt.Println("Err: ", e)
		os.Exit(1)
	}

	result := 0
	fmt.Println("Give answers to the questions")

	bank := formatQandAs(content)

	for i := 0; i < len(bank); i++ {
		pair := bank[i]
		fmt.Println(pair.ques)

		var ans int
		fmt.Scan(&ans)
		fmt.Println(ans, pair.ans)
		if ans == pair.ans {
			result += 1
		}
	}

	fmt.Println("count", result)

}

func formatQandAs(content []byte) []quesAns {

	qa := []quesAns{}

	// get slice of qa
	quesAndAns := strings.Split(string(content), "\n")
	for _, pair := range quesAndAns {
		pairSlice := strings.Split(pair, ",")
		ques := pairSlice[0]
		fmt.Println("hmm", len(pairSlice[1]))
		ans, _ := strconv.Atoi(strings.TrimRight(pairSlice[1], "\n"))
		fmt.Println("hmm", ans)
		p := quesAns{ques: ques, ans: ans}
		qa = append(qa, p)
	}

	return qa

}
