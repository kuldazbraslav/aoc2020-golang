package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func countTrees(lines []string, shiftX int, shiftY int) (cnt int) {
	cnt = 0
	posX := 0
	for posY, r := range lines {
		if posY % shiftY > 0 {
			continue
		}

		if r[posX] == '#' {
			cnt++
		}
		posX = (posX + shiftX) % len(r)
	}
	return cnt
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var cnt [5]int
	cnt[0] = countTrees(lines, 1, 1)
	cnt[1] = countTrees(lines, 3, 1)
	cnt[2] = countTrees(lines, 5, 1)
	cnt[3] = countTrees(lines, 7, 1)
	cnt[4] = countTrees(lines, 1, 2)
	fmt.Println(cnt)

	product := 1
	for _, i := range cnt {
		product *= i
	}
	fmt.Println(product)
}
