package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	programList := make([][]string, 0)
	program := make([]string, 0)
	modelNumber := make([]int, 0)

	for _, inst := range Input_list {

		if checkIfInp(inst) {
			programList = append(programList, program)
			program = make([]string, 0)
		}

		program = append(program, inst)

	}

	programList = append(programList, program)

	if len(programList[0]) == 0 {
		programList = programList[1:]
	}

	varVals := initVars()
	newModelNumber := 0

	for i, val := range programList {

		varVals, newModelNumber = discoveryInpInput(val, varVals)

		modelNumber = append(modelNumber, newModelNumber)

		fmt.Println(varVals)

		if i == 4 {
			break
		}

	}

	//fmt.Println(programList)
	fmt.Println(modelNumber)

}

func discoveryInpInput(program []string, varVals map[string]int) (map[string]int, int) {

	rightNumber := false
	varsMap := make(map[string]int)

	varsMap["w"] = varVals["w"]
	varsMap["x"] = varVals["x"]
	varsMap["y"] = varVals["y"]
	varsMap["z"] = varVals["z"]

	for i := 9; i > 0 && !rightNumber; i-- {

		//fmt.Println(varVals)

		varsMap, rightNumber = runProgram(program, varVals, i)

		//fmt.Println(rightNumber)
		//fmt.Println(varsMap)

		if rightNumber {
			//*varVals = varsMap
			return varsMap, i
		}

	}

	return varVals, 0

}

func runProgram(program []string, varVals map[string]int, inputVal int) (map[string]int, bool) {

	varsMap := make(map[string]int)

	varsMap["w"] = varVals["w"]
	varsMap["x"] = varVals["x"]
	varsMap["y"] = varVals["y"]
	varsMap["z"] = varVals["z"]

	magicNumber := 0
	programType := 1

	for _, inst := range program {

		val, validRes := callFunc(inst, inputVal, varsMap)

		if inst == "div z 1" {

			programType = 1

			//fmt.Println(program[len(program)-3])
			magicNumber, _ = strconv.Atoi(strings.Split(program[len(program)-3], " ")[2])

		} else if inst == "div z 26" {

			programType = 2

			magicNumber, _ = strconv.Atoi(strings.Split(program[len(program)-3], " ")[2])

		}

		if validRes {
			instSplited := strings.Split(inst, " ")
			(varsMap)[instSplited[1]] = val
		} else {
			return varsMap, false
		}

	}

	if varVals["z"]*26+inputVal+magicNumber == varsMap["z"] {
		return varsMap, true
	}

	if varsMap["z"]%26+magicNumber == inputVal {
		return varsMap, true
	}

	if programType == 2 {
		fmt.Println(inputVal, varsMap)
	}

	//fmt.Println(programType)

	return varsMap, (varsMap)["z"] == 0
}

func callFunc(inst string, inputVal int, varVals map[string]int) (int, bool) {

	instSplited := strings.Split(inst, " ")
	switch instSplited[0] {
	case "inp":
		return inputVal, true
	case "add":
		//val1 := instSplited[1]              //(varVals)[instSplited[1]]
		//val2 := instSplited[2]              //(varVals)[instSplited[2]]
		return addVals(instSplited[1], instSplited[2], varVals) //val1 + val2, true
	case "mul":
		//val1 := instSplited[1] //(varVals)[instSplited[1]]
		//val2 := instSplited[2] //(varVals)[instSplited[2]]
		return mulVals(instSplited[1], instSplited[2], varVals)
	case "div":
		//val1 := instSplited[1] //(varVals)[instSplited[1]]
		//val2 := instSplited[2] //(varVals)[instSplited[2]]
		return divVals(instSplited[1], instSplited[2], varVals)
	case "mod":
		//val1 := instSplited[1] //(varVals)[instSplited[1]]
		//val2 := instSplited[2] //(varVals)[instSplited[2]]
		return modVals(instSplited[1], instSplited[2], varVals)
	case "eql":
		//val1 := instSplited[1] //(varVals)[instSplited[1]]
		//val2 := instSplited[2] //(varVals)[instSplited[2]]
		return EqlVals(instSplited[1], instSplited[2], varVals)
	}

	return 0, false

}

func EqlVals(var1, var2 string, varVals map[string]int) (int, bool) {

	val1, _ := (varVals)[var1]
	val2, exist := (varVals)[var2]

	if !exist {
		val2, _ = strconv.Atoi(var2)
	}

	val := 0
	if val1 == val2 {
		val = 1
	}

	return val, true

}

func addVals(var1, var2 string, varVals map[string]int) (int, bool) {

	val1, _ := (varVals)[var1]
	val2, exist := (varVals)[var2]

	if !exist {
		val2, _ = strconv.Atoi(var2)
	}

	return val1 + val2, true

}

func mulVals(var1, var2 string, varVals map[string]int) (int, bool) {

	val1, _ := (varVals)[var1]
	val2, exist := (varVals)[var2]

	if !exist {
		val2, _ = strconv.Atoi(var2)
	}

	return val1 * val2, true

}

func divVals(var1, var2 string, varVals map[string]int) (int, bool) {

	val1, _ := (varVals)[var1]
	val2, exist := (varVals)[var2]

	if !exist {
		val2, _ = strconv.Atoi(var2)
	}

	if val2 == 0 {
		return 0, false
	}

	val := math.Floor(float64(val1) / float64(val2))

	return int(val), true

}

func modVals(var1, var2 string, varVals map[string]int) (int, bool) {

	val1, _ := (varVals)[var1]
	val2, exist := (varVals)[var2]

	if !exist {
		val2, _ = strconv.Atoi(var2)
	}

	if val1 < 0 || val2 <= 0 {
		return 0, false
	}

	val := val1 % val2

	return val, true

}

func initVars() map[string]int {
	valsMap := make(map[string]int)

	valsMap["w"] = 0
	valsMap["x"] = 0
	valsMap["y"] = 0
	valsMap["z"] = 0

	return valsMap
}

func checkIfInp(inst string) bool {

	instName := strings.Split(inst, " ")[0]

	if instName == "inp" {
		return true
	}

	return false

}

var Input_list_t = []string{"inp w",
	"add z w",
	"mod z 2",
	"div w 2",
	"add y w",
	"mod y 2",
	"div w 2",
	"add x w",
	"mod x 2",
	"div w 2",
	"mod w 2"}

var Input_list = []string{
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 10",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 12",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 10",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 10",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 12",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 8",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 11",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 4",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x 0",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 3",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 15",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 10",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 13",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 6",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -12",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 13",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -15",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 8",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -15",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 1",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -4",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 7",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 1",
	"add x 10",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 6",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -5",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 9",
	"mul y x",
	"add z y",
	"inp w",
	"mul x 0",
	"add x z",
	"mod x 26",
	"div z 26",
	"add x -12",
	"eql x w",
	"eql x 0",
	"mul y 0",
	"add y 25",
	"mul y x",
	"add y 1",
	"mul z y",
	"mul y 0",
	"add y w",
	"add y 9",
	"mul y x",
	"add z y"}
