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

func FollowPlannedCourse(input []string) int {
	var horizontal, depth int

	for _, instruction := range input {
		parts := strings.Split(instruction, " ")
		if len(parts) == 2 {
			direction := parts[0]
			amount, err := strconv.Atoi(parts[1])

			mustErr(err)

			switch direction {
			case "down":
				depth = depth + amount

			case "up":
				depth = depth - amount

			case "forward":
				horizontal = horizontal + amount
			}
		}
	}

	return horizontal * depth
}

func FollowPlannedCoursePart2(input []string) int {
	var horizontal, depth, aim int

	for _, instruction := range input {
		parts := strings.Split(instruction, " ")
		if len(parts) == 2 {
			direction := parts[0]
			amount, err := strconv.Atoi(parts[1])

			mustErr(err)

			switch direction {
			case "down":
				aim = aim + amount

			case "up":
				aim = aim - amount

			case "forward":
				horizontal = horizontal + amount
				plus_depth := aim * amount
				depth = depth + plus_depth
			}
		}
	}

	return horizontal * depth
}

func main() {
	input := getInput("input.txt")
	resultPart1 := FollowPlannedCourse(input)
	resultPart2 := FollowPlannedCoursePart2(input)

	fmt.Println(resultPart1)
	fmt.Println(resultPart2)
}
