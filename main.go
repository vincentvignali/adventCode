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

type Elve struct {
	name         int
	totalCalorie int
}
type ElveCollection []Elve

const file = "firstPuzzle.txt"

var (
	str                   = ""
	elves                 = ElveCollection{}
	elveCtr         int64 = 0
	topThreeElveCtr int   = 0
)

func main() {
	// 1. Set up the file and the Scanner
	file, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)

	// 2. Scan the file to build a build string inside str with 'elve' to delimite elves
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {

			str = str + " elve"
		} else {

			str = str + " " + fileScanner.Text()
		}
	}
	// 3. Split the String
	subStrings := strings.Split(str, "elve")

	// 4.1 Iterate over the slice
	for idx, subString := range subStrings {
		// 4.2 Split the string at every space to isolate the numbers
		calorie := strings.Split(subString, " ")
		// 4.3 Remove the first  and the last elements being blank spaces
		calorie = calorie[1 : len(calorie)-1]
		// 4.4 Set this elve counter to 0.
		elveCtr = 0
		// 5.1 Iterate over each number
		for _, v := range calorie {
			// 5.2 Parse it into a number
			parseValue, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				log.Println(err)
			}
			// 5.3 Increase the counter of the current elve
			elveCtr = elveCtr + parseValue
		}
		// 6. Create the elve and append it in the collection.
		elve := Elve{name: idx, totalCalorie: int(elveCtr)}
		elves = append(elves, elve)
	}
	// 7. Sort the collection to have at the first postion the elve with the highest calorie value.
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].totalCalorie > elves[j].totalCalorie
	})
	fmt.Printf("Elve Number %v has a total of  %v calories\n", elves[0].name, elves[0].totalCalorie)

	topThreeElves := elves[:3]
	topThreeElveCtr = 0
	for _, elve := range topThreeElves {
		topThreeElveCtr = topThreeElveCtr + elve.totalCalorie
	}
	fmt.Printf("Top three Number total is %v calories\n", topThreeElveCtr)
}
