package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elve struct {
	name         string
	totalCalorie int
}
type ElveCollection []Elve

const file = "firstPuzzle.txt"

var (
	elves               = ElveCollection{}
	elveCtr         int = 0
	topThreeElveCtr int = 0
	buffer              = []int64{}
)

func main() {
	file, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if len(fileScanner.Text()) != 0 {
			calorie, err := strconv.ParseInt(fileScanner.Text(), 10, 0)
			if err != nil {
				fmt.Println(err)
			}
			buffer = append(buffer, calorie)
		} else {
			var elveTotalCalorie int64 = 0
			for _, v := range buffer {
				elveTotalCalorie += v
			}
			newElve := Elve{name: "Elfe " + strconv.Itoa(elveCtr), totalCalorie: int(elveTotalCalorie)}
			elves = append(elves, newElve)
			buffer = buffer[len(buffer):]
			elveCtr++
		}
	}

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].totalCalorie > elves[j].totalCalorie
	})
	fmt.Printf("Top One Elve: %v ! He has a total of  %v calories.\n", elves[0].name, elves[0].totalCalorie)

	topThreeElves := elves[:3]
	topThreeElveCtr = 0
	for _, elve := range topThreeElves {
		topThreeElveCtr = topThreeElveCtr + elve.totalCalorie
	}
	fmt.Printf("Top three Number total is %v calories.\n", topThreeElveCtr)
}
