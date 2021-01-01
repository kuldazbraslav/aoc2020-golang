package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type recordMap map[rune]int
type record struct {
	answers recordMap
	people int
}

func readRecords(path string) ([]record, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]record, 0, 10)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			i++
			continue
		}

		if len(result) < i+1 {
			var rec record
			rec.answers = make(recordMap)
			result = append(result, rec)
		}
		for _, question := range line {
			result[i].answers[question]++
		}
		result[i].people++
	}
	return result, scanner.Err()
}

func (rec *record) everyoneSaysYesCount() int {
	cnt := 0
	for _, m := range rec.answers {
		if m == rec.people {
			cnt++
		}
	}
	return cnt
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing file path!")
	}
	records, err := readRecords(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	acc := 0
	for _, rec := range records {
		fmt.Println(rec)
		acc += rec.everyoneSaysYesCount()
	}

	fmt.Println(acc)
}
