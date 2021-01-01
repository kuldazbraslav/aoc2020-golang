package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line_t struct {
	limits   [2]int
	reqChar  rune
	password string
}

func parseLine(l string) line_t {
	var result line_t
	lineParts := strings.Split(l, ": ")
	result.password = lineParts[1]
	lineParts = strings.Split(lineParts[0], " ")
	result.reqChar = []rune(lineParts[1])[0]
	lineParts = strings.Split(lineParts[0], "-")
	result.limits[0], _ = strconv.Atoi(lineParts[0])
	result.limits[1], _ = strconv.Atoi(lineParts[1])
	return result
}

func readLines(path string) ([]line_t, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []line_t
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, parseLine(l))
	}
	return lines, scanner.Err()
}

func (r line_t) test() bool {
	return ([]rune(r.password)[r.limits[0]-1] == r.reqChar) != ([]rune(r.password)[r.limits[1]-1] == r.reqChar)
	/*cnt := 0
	for _, c := range r.password {
		if c == r.reqChar {
			cnt++
		}
	}
	return (cnt >= r.limits[0] && cnt <= r.limits[1])*/
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	for _, r := range lines {
		if r.test() {
			cnt++
		}
	}
	fmt.Println(cnt)
}
