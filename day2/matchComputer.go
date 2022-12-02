package main

const (
	loose = 0
	draw  = 3
	win   = 6

	rock     = 1
	paper    = 2
	scissors = 3
)

var (
	matchResult = 0
	myHand      = 0
	gameResult  = 0
)

// Part 1
func computeMatch(opp string, me string) int {
	matchResult = 0
	switch me {
	case "X":
		matchResult += rock + computeMatchX(opp)
	case "Y":
		matchResult += paper + computeMatchY(opp)
	case "Z":
		matchResult += scissors + computeMatchZ(opp)
	}
	return matchResult
}

func computeMatchX(opp string) int {
	switch opp {
	case "A":
		gameResult = draw
	case "B":
		gameResult = loose
	case "C":
		gameResult = win

	}
	return gameResult
}
func computeMatchY(opp string) int {
	switch opp {
	case "A":
		gameResult = win
	case "B":
		gameResult = draw
	case "C":
		gameResult = loose

	}
	return gameResult
}
func computeMatchZ(opp string) int {
	switch opp {
	case "A":
		gameResult = loose
	case "B":
		gameResult = win
	case "C":
		gameResult = draw

	}
	return gameResult
}

// Part 2
func bisComputeMatch(opp string, me string) int {
	matchResult := 0
	switch me {
	case "X":
		matchResult += loose + computeMatchLoose(opp)
	case "Y":
		matchResult += draw + computeMatchDraw(opp)
	case "Z":
		matchResult += win + computeMatchWin(opp)
	}

	return matchResult
}

func computeMatchLoose(opp string) int {
	switch opp {
	case "A":
		myHand = scissors
	case "B":
		myHand = rock
	case "C":
		myHand = paper

	}
	return myHand
}
func computeMatchDraw(opp string) int {
	switch opp {
	case "A":
		myHand = rock
	case "B":
		myHand = paper
	case "C":
		myHand = scissors

	}
	return myHand
}
func computeMatchWin(opp string) int {
	switch opp {
	case "A":
		myHand = paper
	case "B":
		myHand = scissors
	case "C":
		myHand = rock

	}
	return myHand
}
