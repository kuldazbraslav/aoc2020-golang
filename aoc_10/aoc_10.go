package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func loadAdapters(path string) (result []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		result = append(result, num)
	}
	sort.Ints(result)
	result = append(result, result[len(result)-1]+3)
	result = append(result, 0)
	sort.Ints(result)
	return
}

func countDiffs(joltages []int) (counts [3]int) {
	// diffs between joltages
	for i := 1; i < len(joltages); i++ {
		diff := joltages[i] - joltages[i-1]
		counts[diff-1]++
	}
	return
}

func part2(joltages []int) int {
	part2 := 1
	cnt := 0
	for i := 1; i < len(joltages); i++ {
		diff := joltages[i] - joltages[i-1]
		switch diff {
		case 1:
			cnt++
		case 3:
			if (cnt > 0) {
				part2 *= c(cnt)
			}
			cnt = 0
		}
	}
	return part2
}

func min(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func c(i int) int {
	if i == 0 || i == 1 {
		return 1
	}

	sum := 0
	for j := min(3, i); j > 0; j-- {
		sum += c(i-j)
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		panic("Invalid arg count")
	}
	adapters, err := loadAdapters(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(adapters)
	counts := countDiffs(adapters)
	fmt.Println(counts)
	fmt.Println("Part 1:", counts[0]*counts[2])
	fmt.Println("Part 2:", part2(adapters))
}
