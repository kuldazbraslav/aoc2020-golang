package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type recordMap map[string]string

func readRecords(path string) ([]recordMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fieldRegexp := regexp.MustCompile(`(\S+):(\S+)`)
	result := make([]recordMap, 0, 10)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		fields := fieldRegexp.FindAllStringSubmatch(scanner.Text(), -1)
		if len(fields) == 0 {
			i++
			continue
		}

		if len(result) < i+1 {
			result = append(result, make(recordMap))
		}
		for _, field := range fields {
			result[i][field[1]] = field[2]
		}
	}
	return result, scanner.Err()
}

func (record *recordMap) isValid() bool {
	req := map[string]func(string)bool {
		"byr": func(val string) bool {
			i, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return i >= 1920 && i <= 2002
		},
		"iyr": func(val string) bool {
			i, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return i >= 2010 && i <= 2020
		},
		"eyr": func(val string) bool {
			i, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return i >= 2020 && i <= 2030
		},
		"hgt": func(val string) bool {
			hgt := regexp.MustCompile("^([0-9][0-9][0-9]?)(cm|in)$")
			hgtSS := hgt.FindStringSubmatch(val)
			if len(hgtSS) == 0 {
				return false
			}
			hgtVal, _ :=  strconv.Atoi(hgtSS[1])
			return (hgtSS[2] == "cm" && hgtVal >= 150 && hgtVal <= 193) || (hgtVal >= 59 && hgtVal <= 76)
		},
		"hcl": func(val string) bool {
			hcl := regexp.MustCompile("^#[0-9a-f]{6}$")
			return hcl.MatchString(val)
		},
		"ecl": func(val string) bool {
			ecl := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
			return ecl.MatchString(val)
		},
		"pid": func(val string) bool {
			pid := regexp.MustCompile("^[0-9]{9}$")
			return pid.MatchString(val)
		}}
	for field, validate := range req {
		val, exists := (*record)[field]
		if !exists || !validate(val) {
			return false
		}
	}
	return true
}

func main() {
	records, err := readRecords("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validCnt := 0
	for _, rec := range records {
		if rec.isValid() {
			validCnt++
		}
	}

	fmt.Println(validCnt)
}
