package main

import (
	"fmt"
	"strconv"
	"strings"
)

type player struct {
	pos   int
	score int
}

type diceRoll struct {
	nextVal    int
	diceVal    int
	rollNumber int
}

func main() {

	part1()

	//part2()

}

func part2() {

	playerTurn := 1
	//turn := 0

	diceInfo := diceRoll{1, 0, 1}
	listPlayers := make([]player, 0)
	listPlayers = append(listPlayers, player{0, 0})
	listPlayers = append(listPlayers, player{0, 0})

	for _, val := range Input_start {

		playerPos := strings.Split(val, " ")

		playerNumber, _ := strconv.Atoi(playerPos[1])

		listPlayers[playerNumber-1].pos, _ = strconv.Atoi(playerPos[4])

	}

	listWinners := make([]int, 0)
	listWinners = append(listWinners, 0)
	listWinners = append(listWinners, 0)

	playGame(listPlayers, diceInfo, &listWinners, playerTurn)

	fmt.Println(listWinners)

}

func getDiceNextVal(actVal int) int {

	if actVal == 3 {
		return 1
	} else {
		return actVal + 1
	}

}

func playGame(listPlayers []player, diceInfo diceRoll, listWinners *[]int, playerTurn int) {

	endGame := false

	for turn := 0; !endGame; turn++ {

		if diceInfo.rollNumber == 1 {
			diceInfo.diceVal = 0
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 2, 2}, listWinners, playerTurn)
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 3, 2}, listWinners, playerTurn)
			diceInfo.diceVal += 1
			diceInfo.rollNumber = 2
		}

		if diceInfo.rollNumber == 2 {
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 2, 3}, listWinners, playerTurn)
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 3, 3}, listWinners, playerTurn)
			diceInfo.diceVal += 1
			diceInfo.rollNumber = 3
		}

		if diceInfo.rollNumber == 3 {
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 2, 4}, listWinners, playerTurn)
			playGame([]player{player{listPlayers[0].pos, listPlayers[0].score}, player{listPlayers[1].pos, listPlayers[1].score}}, diceRoll{1, diceInfo.diceVal + 3, 4}, listWinners, playerTurn)
			diceInfo.diceVal += 1
		}

		diceInfo.rollNumber = 1
		//fmt.Println(playerTurn, diceInfo.diceVal)

		for movement := 0; movement < diceInfo.diceVal; movement++ {

			listPlayers[playerTurn-1].pos++

			if listPlayers[playerTurn-1].pos > 10 {
				listPlayers[playerTurn-1].pos = 1
			}

		}
		listPlayers[playerTurn-1].score += listPlayers[playerTurn-1].pos

		fmt.Println(playerTurn, listPlayers[playerTurn-1])

		if listPlayers[playerTurn-1].score >= 21 {
			endGame = true
			//playerWinner := getWinner(listPlayers)
			(*listWinners)[playerTurn-1]++
		} else {
			if playerTurn == 1 {
				playerTurn = 2
			} else {
				playerTurn = 1
			}
		}

	}

}

func getWinner(listPlayers []player) int {

	highScore := 0
	playerWinner := 0

	for i, playerStatus := range listPlayers {

		if playerStatus.score > highScore {
			playerWinner = i
			highScore = playerStatus.score
		}

	}

	return playerWinner

}

func part1() {
	player1 := player{0, 0}
	endgame := false
	diceVal := 1

	listPlayers := make([]player, 0)
	listPlayers = append(listPlayers, player1)
	listPlayers = append(listPlayers, player1)

	for _, val := range Input_start {

		playerPos := strings.Split(val, " ")

		playerNumber, _ := strconv.Atoi(playerPos[1])

		listPlayers[playerNumber-1].pos, _ = strconv.Atoi(playerPos[4])

	}

	playerTurn := 1
	turn := 0
	diceRoll := 0
	for turn = 0; !endgame; turn++ {

		diceRoll += 3

		movs := diceVal
		if diceVal == 100 {
			diceVal = 1
		} else {
			diceVal++
		}
		movs += diceVal
		if diceVal == 100 {
			diceVal = 1
		} else {
			diceVal++
		}
		movs += diceVal

		if diceVal == 100 {
			diceVal = 1
		} else {
			diceVal++
		}

		//fmt.Println(movs)

		for movement := 0; movement < movs; movement++ {

			listPlayers[playerTurn-1].pos++

			if listPlayers[playerTurn-1].pos > 10 {
				listPlayers[playerTurn-1].pos = 1
			}

		}
		listPlayers[playerTurn-1].score += listPlayers[playerTurn-1].pos

		if listPlayers[playerTurn-1].score >= 1000 {
			endgame = true
		} else {
			if playerTurn == 1 {
				playerTurn = 2
			} else {
				playerTurn = 1
			}
		}

	}

	fmt.Println(diceVal)

	fmt.Println(listPlayers)

	fmt.Println(diceRoll)

	lowScore := getLowScore(listPlayers)

	fmt.Println(diceRoll * lowScore)
}

func getLowScore(listPlayers []player) int {

	if len(listPlayers) == 0 {
		return 0
	}

	lowScore := listPlayers[0].score

	for _, player := range listPlayers {

		pScore := player.score

		if lowScore > pScore {
			lowScore = pScore
		}

	}

	return lowScore

}

var Input_start = []string{"Player 1 starting position: 4",
	"Player 2 starting position: 8"}

var Input_start_f = []string{
	"Player 1 starting position: 4",
	"Player 2 starting position: 9"}
