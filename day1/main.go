package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func mustErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getInput(filename string) []string {
	file_content, err := ioutil.ReadFile(filename)

	mustErr(err)

	return strings.Split(string(file_content), "\n")
}

func countMessurementIncrease(input []string) int {
	last := 0
	increase_count := -1

	for _, str_num := range input {
		if str_num == "" {
			continue
		}
		s, err := strconv.Atoi(str_num)
		mustErr(err)
		if int(s) > last {
			increase_count++
		}
		last = int(s)
	}
	return increase_count
}

func countTripleMessurementIncrease(input []string) int {
	last := 0
	increase_count := -1

	for i, str_num := range input {
		if str_num == "" {
			continue
		}
		res := input[i:]
		if len(res) >= 3 {
			a, _ := strconv.Atoi(res[0])
			b, _ := strconv.Atoi(res[1])
			c, _ := strconv.Atoi(res[2])

			if (a + b + c) > last {
				increase_count++
			}
			last = a + b + c
		}
	}

	return increase_count
}

func main() {
	content := getInput("input.txt")

	fmt.Println(countMessurementIncrease(content))
	fmt.Println(countTripleMessurementIncrease(content))
}
