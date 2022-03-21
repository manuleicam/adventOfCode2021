package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkIfExist(arr []int, val int) bool {

	for _, arrVal := range arr {

		for val == arrVal {
			return true
		}

	}
	return false

}

func main() {

	instSplit := strings.Split(Input_list, ": ")
	coordsPair := strings.Split(instSplit[1], ", ")
	coordsPairX := strings.Split(coordsPair[0], "..")
	coordsPairY := strings.Split(coordsPair[1], "..")

	coordsPairX1, _ := strconv.Atoi(strings.Split(coordsPairX[0], "=")[1])
	coordsPairX2, _ := strconv.Atoi(coordsPairX[1])
	coordsPairY1, _ := strconv.Atoi(strings.Split(coordsPairY[0], "=")[1])
	coordsPairY2, _ := strconv.Atoi(coordsPairY[1])

	fmt.Println(coordsPairX1, coordsPairX2, coordsPairY1, coordsPairY2)
	minX, lastV := 0, 0
	minsPoss := make([]int, 0)
	for minX = 0; minX < coordsPairX2; minX++ {

		lastV += minX

		if lastV >= coordsPairX1 && lastV <= coordsPairX2 {
			minsPoss = append(minsPoss, minX)
		}

	}

	minsPoss2 := seqVals(minsPoss[0], coordsPairX2)

	velX := 0
	velY := 0
	//searching := true

	maxHitTarget := -1 //simThrow(velX, velY, coordsPairX1, coordsPairX2, coordsPairY1, coordsPairY2)
	hitCount := 0

	for _, val := range minsPoss2 {
		velX = val
		for velY = 0; velY < (absCos(coordsPairY1) + absCos(coordsPairY1) - absCos(coordsPairY2)); velY++ {

			if !checkIfExist(minsPoss2, val) {
				continue
			}

			newHigh := simThrow(velX, velY, coordsPairX1, coordsPairX2, coordsPairY1, coordsPairY2)

			if newHigh.hitFlag {
				if newHigh.maxHight > maxHitTarget {
					maxHitTarget = newHigh.maxHight
				}
				hitCount++
			} else {
				//searching = false
			}

		}

		for velY = -1; velY >= coordsPairY1; velY-- {

			newHigh := simThrow(velX, velY, coordsPairX1, coordsPairX2, coordsPairY1, coordsPairY2)

			if newHigh.hitFlag {
				if newHigh.maxHight > maxHitTarget {
					maxHitTarget = newHigh.maxHight
				}
				hitCount++
			} else {
				//searching = false
			}

		}

	}

	fmt.Println(maxHitTarget)
	fmt.Printf("Acertou: %v vezes\n", hitCount)

}

func seqVals(ini, end int) []int {

	newArr := make([]int, 0)

	for i := 0; i <= (end - ini); i++ {

		newArr = append(newArr, ini+i)

	}

	return newArr

}

type targetHit struct {
	maxHight int
	hitFlag  bool
}

func simThrow(velX, velY int, x1, x2, y1, y2 int) targetHit {
	posX := 0
	posY := 0
	maxY := 0

	for i := 0; i < 200000; i++ {

		if checkIfTarget(posX, posY, x1, x2, y1, y2) {
			//fmt.Printf("Altura máxima: %v\n", maxY)
			//fmt.Printf("X Quando atingiu: %v\n", posX)
			return targetHit{maxY, true}
		}

		if posY > maxY {
			maxY = posY
		}

		posX += velX
		posY += velY

		if velX < 0 {
			velX++
		} else if velX > 0 {
			velX--
		}
		velY--

	}

	return targetHit{0, false}

}

func minPossX(velX, targetX1, targetX2 int) bool {

	lastX := (velX / 2.0) * (1 + velX)

	fmt.Println(lastX, velX)

	return (lastX >= targetX1 && lastX <= targetX2)

}

func checkIfTarget(coordX, coordY int, x1, x2, y1, y2 int) bool {

	if coordX >= x1 && coordX <= x2 && coordY >= y1 && coordY <= y2 {
		return true
	}

	return false

}

func absCos(num int) int {

	if num < 0 {
		num *= -1
	}

	return num

}

var Input_list_teste = "target area: x=20..30, y=-10..-5"

var Input_list = "target area: x=235..259, y=-118..-62"
