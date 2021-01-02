package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type xmas struct {
	data     []int
	preamble int
}

func loadXmas(path string, preamble int) (result xmas, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		result.data = append(result.data, num)
	}
	result.preamble = preamble
	return
}

func (cipher *xmas) isIndexValid(idx int) bool {
	if idx < cipher.preamble {
		log.Printf("idx %d is less than preamble %d", idx, cipher.preamble)
		return false
	}
	for i, x := range cipher.data[idx-cipher.preamble : idx] {
		for _, y := range cipher.data[idx-cipher.preamble+i+1 : idx] {
			if x+y == cipher.data[idx] {
				return true
			}
		}
	}
	return false
}

func (cipher *xmas) findInvalid() int {
	for i, x := range cipher.data[cipher.preamble:] {
		if !cipher.isIndexValid(i + cipher.preamble) {
			return x
		}
	}
	return -1
}

func minMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
			if max < value {
					max = value
			}
			if min > value {
					min = value
			}
	}
	return min, max
}

func main() {
	if len(os.Args) < 3 {
		panic("Invalid arg count")
	}
	preamble, err := strconv.Atoi(os.Args[2])
	cipher, err := loadXmas(os.Args[1], preamble)
	if err != nil {
		panic(err)
	}
	invalid := cipher.findInvalid()
	fmt.Println("Part 1:", invalid)
}
