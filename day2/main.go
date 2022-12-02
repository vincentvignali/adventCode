package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	totalPointPartOne = 0
	totalPointPartTwo = 0
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		totalPointPartOne += (computeMatch(input[0], input[1]))
		totalPointPartTwo += (bisComputeMatch(input[0], input[1]))

	}
	fmt.Println("TotalPoint - PartOne : ", totalPointPartOne)
	fmt.Println("TotalPoint - Part Two : ", totalPointPartTwo)

}
