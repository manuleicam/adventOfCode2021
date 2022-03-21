package main

import (
	"fmt"
)

func main() {

	populationsBirthDate := [9]int{}

	for _, val := range Input_List {

		populationsBirthDate[val]++

	}

	fmt.Println(populationsBirthDate)

	for i := 0; i < 256; i++ {

		shiftLeftArray(&populationsBirthDate)

	}

	fmt.Println(sumArray(populationsBirthDate))

	//for _, val := range Input_List {
	//	nSons += calcNumberSons(lifeSpan, val, 0)
	//}

}

func sumArray(arrToSum [9]int) int {
	tots := 0

	for _, val := range arrToSum {
		tots += val
	}

	return tots

}

func shiftLeftArray(arrToShift *[9]int) {

	aux := (*arrToShift)[0]
	//(*arrToShift)[7] = (*arrToShift)[8]

	for i := 0; i < 8; i++ {

		(*arrToShift)[i] = (*arrToShift)[i+1]

	}

	(*arrToShift)[6] += aux

	(*arrToShift)[8] = aux

}

func checkForNewFish(fishList *[]int) int {

	numberFishs := 0
	newFishs := 0

	for i, fish := range *fishList {

		if fish == 0 {

			(*fishList)[i] = 6

			newFishs++
		} else {
			(*fishList)[i]--
		}
		numberFishs++
	}

	for i := 0; i < newFishs; i++ {
		*fishList = append(*fishList, 8)
		numberFishs++
	}

	return numberFishs

}

//var Input_List = []int{3, 4, 3, 1, 2}

var Input_List = []int{4, 5, 3, 2, 3, 3, 2, 4, 2, 1, 2, 4, 5, 2, 2, 2, 4, 1, 1, 1, 5, 1, 1, 2, 5, 2, 1, 1, 4, 4, 5, 5, 1, 2, 1, 1, 5, 3, 5, 2, 4, 3, 2, 4, 5, 3, 2, 1, 4, 1, 3, 1, 2, 4, 1, 1, 4, 1, 4, 2, 5, 1, 4, 3, 5, 2, 4, 5, 4, 2, 2, 5, 1, 1, 2, 4, 1, 4, 4, 1, 1, 3, 1, 2, 3, 2, 5, 5, 1, 1, 5, 2, 4, 2, 2, 4, 1, 1, 1, 4, 2, 2, 3, 1, 2, 4, 5, 4, 5, 4, 2, 3, 1, 4, 1, 3, 1, 2, 3, 3, 2, 4, 3, 3, 3, 1, 4, 2, 3, 4, 2, 1, 5, 4, 2, 4, 4, 3, 2, 1, 5, 3, 1, 4, 1, 1, 5, 4, 2, 4, 2, 2, 4, 4, 4, 1, 4, 2, 4, 1, 1, 3, 5, 1, 5, 5, 1, 3, 2, 2, 3, 5, 3, 1, 1, 4, 4, 1, 3, 3, 3, 5, 1, 1, 2, 5, 5, 5, 2, 4, 1, 5, 1, 2, 1, 1, 1, 4, 3, 1, 5, 2, 3, 1, 3, 1, 4, 1, 3, 5, 4, 5, 1, 3, 4, 2, 1, 5, 1, 3, 4, 5, 5, 2, 1, 2, 1, 1, 1, 4, 3, 1, 4, 2, 3, 1, 3, 5, 1, 4, 5, 3, 1, 3, 3, 2, 2, 1, 5, 5, 4, 3, 2, 1, 5, 1, 3, 1, 3, 5, 1, 1, 2, 1, 1, 1, 5, 2, 1, 1, 3, 2, 1, 5, 5, 5, 1, 1, 5, 1, 4, 1, 5, 4, 2, 4, 5, 2, 4, 3, 2, 5, 4, 1, 1, 2, 4, 3, 2, 1}

// {3, 4, 3, 1, 2} -> {7,8,7,}
// 3 -> 1 filho de 6 em 6 dias + 1 filho ao fim de 4 dias
// 4 -> 1 filho de 6 em 6 dias + 1 filho ao fim de 5 dias

// 3 -> (18-3)/6 + 1 -> floor(3.5)    -> 3
// 8 -> (18-8-3)/6 -> floor(1.17)  	  -> 1
// 8 -> (18-(8*2)-3/6) -> floor(0.16) -> 0
