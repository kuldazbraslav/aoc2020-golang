package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func seatToID(seat string) int {
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	res, _ := strconv.ParseInt(r.Replace(seat), 2, 0)
	return int(res)
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	seatIDs := make([]int, len(lines), len(lines))
	highest := 0
	for i, seat := range lines {
		seatIDs[i] = seatToID(seat)
		if seatIDs[i] > highest {
			highest = seatIDs[i]
		}
	}
	fmt.Println(highest)

	sort.Ints(seatIDs)
	offset := seatIDs[0]
	for i, seat := range seatIDs {
		if (i + offset != seat) {
			fmt.Println(seat - 1)
			break
		}
	}
}
