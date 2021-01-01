package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]uint, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, uint(num))
	}
	return lines, scanner.Err()
}

func main() {
	nums, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	/*nums := []uint{1721,
	979,
	366,
	299,
	675,
	1456}
	*/

	for i, n := range nums {
		// fmt.Println(n)
		for j, m := range nums[i+1:] {
			// fmt.Println(n+m)
			for _, l := range nums[j+1:] {
				if n+m+l == 2020 {
                    fmt.Println(n)
                    fmt.Println(m)
                    fmt.Println(l)
					fmt.Println(n * m * l)
					return
				}
			}
		}
	}
}
