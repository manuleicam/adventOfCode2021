package main

import (
	"fmt"
	"strconv"
)

func main() {

	//fmt.Println("sss")

	inputMatrix := make([][]int, 10)

	for i, val := range InputList {

		row := make([]int, 0)

		for _, number := range val {

			newElem, _ := strconv.Atoi(string(number))

			row = append(row, newElem)

		}

		inputMatrix[i] = row

	}
	nFlashs := 0

	nElesm := len(inputMatrix) * len(inputMatrix[1])

	fmt.Println(nElesm)

	for i := 0; nFlashs != 100; i++ {
		nFlashs = 0

		//fmt.Println("STEP---->", i)

		for y, row := range inputMatrix {

			for x, _ := range row {

				incNumber(&inputMatrix, x, y)

			}

		}

		for y, row := range inputMatrix {

			for x, number := range row {

				if number > 9 {
					inputMatrix[y][x] = 0
					nFlashs++
				}

			}

		}
		//fmt.Println(inputMatrix)
		fmt.Println(i)
		fmt.Println(nFlashs)
	}

}

func catchOutBoundsError() {
	if r := recover(); r != nil {
		//return 0
	}
}

func incNumber(numberMatrix *[][]int, x int, y int) {

	defer catchOutBoundsError()

	//if (*numberMatrix)[y][x] < 9 {
	(*numberMatrix)[y][x]++

	if (*numberMatrix)[y][x] == 10 {
		flashLight(numberMatrix, x, y)
	}
	//}

}

func flashLight(numberMatrix *[][]int, x int, y int) {

	//incNumber(numberMatrix, x, y)
	incNumber(numberMatrix, x+1, y)
	incNumber(numberMatrix, x-1, y)
	incNumber(numberMatrix, x, y+1)
	incNumber(numberMatrix, x, y-1)
	incNumber(numberMatrix, x+1, y+1)
	incNumber(numberMatrix, x+1, y-1)
	incNumber(numberMatrix, x-1, y+1)
	incNumber(numberMatrix, x-1, y-1)

}

var InputList_teste = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526"}

var InputList = []string{
	"1224346384",
	"5621128587",
	"6388426546",
	"1556247756",
	"1451811573",
	"1832388122",
	"2748545647",
	"2582877432",
	"3185643871",
	"2224876627"}
