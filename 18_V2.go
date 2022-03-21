package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	part1()

	val := 0

	EqList := Input_list //Input_list_Split //

	for _, line := range EqList {

		for _, line2 := range EqList {

			valAux := part2(line, line2)

			if valAux > val {
				val = valAux
			}

		}

	}
	fmt.Println(val)

}

func part2(eqEsq string, eqDir string) int {

	eqList := make([]string, 0)
	eqList = append(eqList, eqEsq)
	eqList = append(eqList, eqDir)

	nParents := make([]int, 0)
	eqArr := make([]int, 0)

	for i, line := range eqList {

		nParentsAux, eqArrAux := buidEQ(line, i)

		//fmt.Println(nParents, "+", nParentsAux)
		//fmt.Println(eqArr, "+", eqArrAux)

		sumEQ(&nParents)

		nParents = append(nParents, nParentsAux...)

		eqArr = append(eqArr, eqArrAux...)

		//fmt.Println(nParents)
		//fmt.Println(eqArr)

		checkRules(&nParents, &eqArr)

		//if i == 16 {
		//	break
		//}

	}

	//fmt.Println(nParents)
	//fmt.Println(eqArr)

	for i := 0; i < len(eqArr)-1; i++ {

		if (nParents)[i] == (nParents)[i+1] {

			valEsq := (eqArr)[i] * 3
			valDir := (eqArr)[i+1] * 2

			(eqArr)[i] = valEsq + valDir
			nParents[i]--

			eqArr = append(eqArr[:i+1], eqArr[i+2:]...)
			nParents = append(nParents[:i+1], nParents[i+2:]...)

			i = -1

		}

	}

	// 3702 -> To High
	//fmt.Println(nParents)
	//fmt.Println(eqArr)
	return eqArr[0]
}

func part1() {
	nParents := make([]int, 0)
	eqArr := make([]int, 0)

	for i, line := range Input_list {

		nParentsAux, eqArrAux := buidEQ(line, i)

		//fmt.Println(nParents, "+", nParentsAux)
		//fmt.Println(eqArr, "+", eqArrAux)

		sumEQ(&nParents)

		nParents = append(nParents, nParentsAux...)

		eqArr = append(eqArr, eqArrAux...)

		//fmt.Println(nParents)
		//fmt.Println(eqArr)

		checkRules(&nParents, &eqArr)

		//if i == 16 {
		//	break
		//}

	}

	//fmt.Println(nParents)
	//fmt.Println(eqArr)

	for i := 0; i < len(eqArr)-1; i++ {

		if (nParents)[i] == (nParents)[i+1] {

			valEsq := (eqArr)[i] * 3
			valDir := (eqArr)[i+1] * 2

			(eqArr)[i] = valEsq + valDir
			nParents[i]--

			eqArr = append(eqArr[:i+1], eqArr[i+2:]...)
			nParents = append(nParents[:i+1], nParents[i+2:]...)

			i = -1

		}

	}

	// 3702 -> To High
	//fmt.Println(nParents)
	fmt.Println(eqArr)
}

func checkRules(nParents, eqArr *[]int) {

	eqSize := len(*nParents)
	//x := true
	//for x {
	nOcores := 0
	for i := 0; i < eqSize; i++ {

		old_i := i

		if (*nParents)[old_i] > 4 {

			//fmt.Println("BEFORE EXPLODE", *eqArr, " ", *nParents)
			ApplyExplode(nParents, eqArr, old_i, eqSize)

			nOcores++
			//fmt.Println("AFTER EXPLODE", *eqArr, " ", *nParents)
			eqSize = len(*nParents)
			//i = 0
			//continue
		}

		eqSize = len(*nParents)

	}

	for i := 0; i < eqSize; i++ {
		if (*eqArr)[i] > 9 {

			//fmt.Println("BEFORE SPLIT", *eqArr, " ", *nParents)

			ApplySplit(nParents, eqArr, i, eqSize)

			//fmt.Println("AFTER SPLIT", *eqArr, " ", *nParents)

			eqSize = len(*nParents)

			if (*nParents)[i] > 4 {

				//fmt.Println("BEFORE 2º EXPLODE", *eqArr, " ", *nParents)

				ApplyExplode(nParents, eqArr, i, eqSize)

				//fmt.Println("AFTER 2º EXPLODE", *eqArr, " ", *nParents)

			}
			i = -1
			nOcores++

			eqSize = len(*nParents)
		}

	}

	//if nOcores == 0 {
	//	x = false
	//}

	//}

	//fmt.Println(*nParents)
	//fmt.Println(*eqArr)

}

func addArrayVal(pos, lenArr int, eqArr *[]int) {

	lastval := (*eqArr)[pos]

	for i := pos + 1; i < lenArr; i++ {
		aux := (*eqArr)[i]
		(*eqArr)[i] = lastval
		lastval = aux

	}

	(*eqArr) = append((*eqArr), lastval)

}

func ApplySplit(nParents, eqArr *[]int, pos, eqSize int) {

	valToSplit := (*eqArr)[pos]
	valEsq := valToSplit / 2
	valDir := int(math.Ceil(float64(valToSplit) / 2))
	nParentsSplit := (*nParents)[pos]

	if pos < eqSize {

		(*eqArr)[pos] = valEsq
		addArrayVal(pos, eqSize, eqArr)
		(*eqArr)[pos+1] = valDir
		//eqArrAux := (*eqArr)[pos+1:]
		//fmt.Println("WWWW", eqArrAux)
		//(*eqArr) = append((*eqArr)[:pos+1], valDir)
		//fmt.Println("WWWW", eqArrAux)
		//(*eqArr) = append((*eqArr), eqArrAux...)

		(*nParents)[pos] = nParentsSplit + 1
		addArrayVal(pos, eqSize, nParents)
		(*nParents)[pos+1] = nParentsSplit + 1
		//nParentsAux := (*nParents)[pos+1:]
		//(*nParents) = append((*nParents)[:pos+1], nParentsSplit+1)
		//(*nParents) = append((*nParents), nParentsAux...)

	} else {
		(*eqArr) = append((*eqArr), valDir)
		(*eqArr)[pos] = valEsq
		(*nParents) = append((*nParents), nParentsSplit+1)
		(*nParents)[pos] = nParentsSplit + 1
	}

}

func ApplyExplode(nParents, eqArr *[]int, pos, eqSize int) {

	if pos > 0 {
		(*eqArr)[pos-1] += (*eqArr)[pos]
	}
	(*nParents)[pos]--
	(*eqArr)[pos] = 0

	if pos < eqSize-2 {
		(*eqArr)[pos+2] += (*eqArr)[pos+1]

		(*eqArr) = append((*eqArr)[:pos+1], (*eqArr)[pos+2:]...)
		(*nParents) = append((*nParents)[:pos+1], (*nParents)[pos+2:]...)
	} else {

		(*eqArr) = (*eqArr)[:pos+1]
		(*nParents) = (*nParents)[:pos+1]
	}

}

func sumEQ(nParents *[]int) {

	for i, _ := range *nParents {
		(*nParents)[i]++
	}

}

func buidEQ(line string, linePos int) ([]int, []int) {
	parsNumber := 0
	if linePos == 0 {
		parsNumber--
	}

	nParents := make([]int, 0)
	eqArr := make([]int, 0)
	for _, char := range line {

		if string(char) == "[" {
			parsNumber++
		} else if string(char) == "]" {
			parsNumber--
		} else if val, err := strconv.Atoi(string(char)); err == nil {
			nParents = append(nParents, parsNumber+1)
			eqArr = append(eqArr, val)
		}

	}

	return nParents, eqArr

}

var Input_test_res = []string{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"}

var Input_list_te = []string{
	"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
	"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
	"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
	"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
	"[7,[5,[[3,8],[1,4]]]]",
	"[[2,[2,2]],[8,[8,1]]]",
	"[2,9]",
	"[1,[[[9,3],9],[[9,0],[0,7]]]]",
	"[[[5,[7,4]],7],1]",
	"[[[[4,2],2],6],[8,7]]"}

var Input_list_Exp = []string{
	"[1,1]",
	"[2,2]",
	"[3,3]",
	"[4,4]",
	"[5,5]",
	"[6,6]"}

var Input_list_Split = []string{
	"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
	"[[[5,[2,8]],4],[5,[[9,9],0]]]",
	"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
	"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
	"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
	"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
	"[[[[5,4],[7,7]],8],[[8,3],8]]",
	"[[9,3],[[9,9],[6,[4,9]]]]",
	"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
	"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"}

var Input_list = []string{
	"[[[2,[3,5]],[8,7]],[[9,3],2]]",
	"[[3,[3,7]],[[3,6],[[1,1],7]]]",
	"[8,[[5,5],[2,9]]]",
	"[[5,[3,5]],[[2,1],[[7,1],[7,7]]]]",
	"[[[[3,3],0],[[0,3],0]],[[8,[2,2]],[[0,4],3]]]",
	"[3,6]",
	"[[5,[[4,2],1]],[[6,[0,3]],[4,[7,7]]]]",
	"[[6,5],[2,[3,6]]]",
	"[[[[0,1],0],[[7,4],5]],[[6,2],[4,[0,8]]]]",
	"[[[[4,7],3],8],[[7,[0,4]],[7,[1,4]]]]",
	"[[[0,[9,8]],[2,9]],[[[6,4],[4,0]],4]]",
	"[2,[[4,[8,5]],[6,8]]]",
	"[[[0,7],[5,[3,0]]],[[[6,4],[3,2]],[[4,7],[9,6]]]]",
	"[[[[0,6],[0,7]],[8,0]],[8,[4,8]]]",
	"[[[[9,9],2],[[6,2],[2,2]]],[[5,[8,8]],6]]",
	"[[0,[[4,6],7]],[[7,[4,8]],9]]",
	"[[0,5],[[5,3],[[3,9],4]]]",
	"[2,[[[9,4],[8,8]],1]]",
	"[5,[[[2,3],6],[2,[7,0]]]]",
	"[[7,[[8,6],3]],[2,[2,7]]]",
	"[6,[[2,4],[[9,7],[5,9]]]]",
	"[[[9,[2,1]],9],1]",
	"[[[6,9],[2,[2,5]]],[[[4,4],0],7]]",
	"[1,[[[3,9],[6,1]],[4,0]]]",
	"[[[3,8],[3,[2,7]]],[[[9,2],2],6]]",
	"[6,[[8,[3,1]],7]]",
	"[[[9,9],7],[[[3,1],[8,4]],[0,0]]]",
	"[[[1,[7,8]],[4,2]],2]",
	"[[9,7],6]",
	"[[6,[4,8]],[[[8,6],[0,1]],[[0,4],[8,4]]]]",
	"[[[[1,8],[8,6]],[9,[2,0]]],[5,[2,[7,2]]]]",
	"[1,9]",
	"[[8,[9,[9,3]]],[[[1,1],8],[[1,5],[8,6]]]]",
	"[[[3,[4,4]],3],[[7,0],[6,0]]]",
	"[[[6,[6,3]],[6,7]],[1,[8,0]]]",
	"[[[9,7],[1,7]],8]",
	"[[8,[[4,6],[4,8]]],8]",
	"[[[1,9],6],1]",
	"[[[[0,5],[0,0]],7],[4,8]]",
	"[[[[6,0],[4,2]],[8,[5,1]]],[[0,[4,8]],[[3,2],8]]]",
	"[[[[5,9],[5,8]],[9,[0,1]]],[[[8,6],[3,1]],[[9,8],0]]]",
	"[0,[[9,9],[6,2]]]",
	"[[[[7,9],[9,1]],[[1,0],[6,4]]],[4,[[2,1],2]]]",
	"[4,2]",
	"[[[6,5],[[0,6],2]],[[[1,2],0],[[8,9],8]]]",
	"[[8,[[4,1],0]],[[[1,5],[3,5]],3]]",
	"[[[8,3],[[9,1],[8,1]]],[[9,9],3]]",
	"[[2,7],[[[3,9],[2,3]],9]]",
	"[[2,[[7,3],[1,6]]],[[4,4],[2,7]]]",
	"[[[5,6],[3,[5,3]]],[[[2,8],0],[4,[8,8]]]]",
	"[[[1,2],[4,[5,8]]],[8,[8,[9,0]]]]",
	"[[[[0,5],[8,1]],0],[[[5,4],[6,9]],[[7,5],[4,9]]]]",
	"[[9,[2,1]],[[[3,8],[9,5]],[[4,4],4]]]",
	"[[[5,9],[[1,1],[8,9]]],[[1,9],8]]",
	"[[[8,8],[3,9]],[[[2,1],0],9]]",
	"[[[[7,8],2],[5,[3,9]]],[6,1]]",
	"[[[[2,4],[9,1]],[[9,8],[4,4]]],[0,1]]",
	"[[[[8,8],0],9],4]",
	"[[[8,[1,5]],0],[[[8,5],4],[[7,3],[9,5]]]]",
	"[[[5,4],[[5,1],2]],[[[6,8],6],[[3,6],[1,9]]]]",
	"[[[3,[2,5]],[6,[6,2]]],[[0,7],[3,9]]]",
	"[3,[[2,9],8]]",
	"[[[[3,7],[1,6]],[[9,9],[0,3]]],[[[7,3],8],[[3,1],6]]]",
	"[[[[7,1],4],[[4,0],[4,5]]],[8,[[5,3],[4,6]]]]",
	"[[[[0,8],1],[7,9]],[[7,5],[[1,0],[0,9]]]]",
	"[[[9,7],[0,[7,8]]],2]",
	"[[[5,2],5],[0,[[1,6],[2,0]]]]",
	"[[[[3,9],7],7],[[3,[3,4]],[0,[5,9]]]]",
	"[[[[2,5],[9,9]],[1,[6,5]]],6]",
	"[[[1,[5,9]],[[1,1],1]],[5,[[0,4],[9,0]]]]",
	"[[[5,8],[0,7]],[3,[2,[8,6]]]]",
	"[[[[0,7],[7,9]],[[8,4],[8,7]]],[0,[[3,7],9]]]",
	"[[[5,[5,5]],[[9,5],8]],[[[2,1],5],9]]",
	"[5,[4,[[3,6],[3,2]]]]",
	"[[[9,4],3],[[[8,7],[7,5]],[8,[7,7]]]]",
	"[9,[[[9,2],0],[[9,9],[4,3]]]]",
	"[[[4,[7,2]],[[7,9],[5,4]]],1]",
	"[[[[4,9],5],7],[[5,6],0]]",
	"[[[5,[3,1]],[8,1]],[8,[7,0]]]",
	"[[5,6],[6,[[0,5],0]]]",
	"[[[5,[4,5]],9],6]",
	"[[[9,[7,0]],6],[2,[1,6]]]",
	"[[[9,[8,4]],[7,[6,0]]],[[[4,6],[7,5]],[8,[0,8]]]]",
	"[0,7]",
	"[[3,[3,8]],[9,[[3,1],[4,4]]]]",
	"[[6,7],[8,9]]",
	"[[[[9,8],[0,2]],[[4,0],[7,5]]],[[[5,0],1],2]]",
	"[[[[1,2],[3,9]],1],[[5,1],[0,1]]]",
	"[[[[5,8],0],6],[7,0]]",
	"[[[8,[5,4]],[[3,0],7]],[[8,[7,5]],4]]",
	"[[[[5,8],8],8],[[[0,4],[2,5]],0]]",
	"[[[9,6],3],[[[3,3],1],[2,[9,2]]]]",
	"[[[6,3],6],[[[4,1],8],[2,3]]]",
	"[2,[[1,8],0]]",
	"[5,[[[7,6],[1,9]],[4,[8,2]]]]",
	"[[[[6,9],[0,7]],[[2,7],8]],[[6,0],[2,[1,6]]]]",
	"[[[[7,8],[5,1]],[[2,9],2]],0]",
	"[5,3]",
	"[2,[7,[7,[5,8]]]]",
	"[[3,3],[8,[2,6]]]"}
