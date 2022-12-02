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

const (
	file         = "firstPuzzle.txt"
	topRange int = 3
)

var (
	elves         = ElveCollection{}
	elveCtr int   = 0
	acc     int64 = 0
	topCtr  int   = 0
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
			acc += calorie
		} else {
			newElve := Elve{name: "Elfe " + fmt.Sprintf("%v", elveCtr), totalCalorie: int(acc)}
			elves = append(elves, newElve)
			elveCtr++
			acc = 0
		}
	}

	// Part 1
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].totalCalorie > elves[j].totalCalorie
	})
	fmt.Printf("Top One Elve: %v ! He has a total of  %v calories.\n", elves[0].name, elves[0].totalCalorie)

	// Part 2
	t := elves[:topRange]
	for _, elve := range t {
		topCtr += elve.totalCalorie
	}
	fmt.Printf("Top three Number total is %v calories.\n", topCtr)
}
