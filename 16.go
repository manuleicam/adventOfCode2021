package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	mapHexBit := make(map[string]string)
	bitTrans := make([]string, 0)

	for _, val := range map_HexaToBin {

		instSplit := strings.Split(val, " = ")

		mapHexBit[instSplit[0]] = instSplit[1]

	}
	fmt.Println(mapHexBit)

	for _, char := range Input_pack {

		bitChar := mapHexBit[string(char)]

		for _, bit := range bitChar {

			bitTrans = append(bitTrans, string(bit))

		}

	}

	//fmt.Println(bitTrans)

	//cleanTransmission(&bitTrans)

	//fmt.Println(bitTrans)

	versionTot := 0
	expr := make([]string, 0)
	decodePackage(&bitTrans, &versionTot, &expr)

	fmt.Println(versionTot)
	fmt.Println(expr)

	op := pop(&expr)

	totEq := handleOp(op, &expr)

	fmt.Println(totEq)

}

func convertTypeId(typeId int) string {

	switch typeId {

	case 0:
		return "+"
	case 1:
		return "*"
	case 2:
		return "min"
	case 3:
		return "max"
	case 5:
		return ">"
	case 6:
		return "<"
	case 7:
		return "="

	}

	return ""

}

func decodePackage(bitTrans *[]string, versionTot *int, expr *[]string) {

	if len(*bitTrans) < 6 {
		*bitTrans = make([]string, 0)
		return
	}

	version := calcBin((*bitTrans)[0] + (*bitTrans)[1] + (*bitTrans)[2])
	typeId := calcBin((*bitTrans)[3] + (*bitTrans)[4] + (*bitTrans)[5])

	num := 0

	*versionTot += version

	fmt.Printf("--------VERSION: %v\n", version)

	(*bitTrans) = (*bitTrans)[6:]

	switch typeId {
	case 4:
		//fmt.Println("isNumber")
		num = isNumber(bitTrans)
		*expr = append(*expr, strconv.Itoa(num))
		fmt.Printf("NUMBER: %v\n", num)
	default:
		//fmt.Println("isOperator")
		opString := convertTypeId(typeId)
		*expr = append(*expr, opString)
		*expr = append(*expr, "(")
		isOperator(bitTrans, versionTot, expr)
	}

	//fmt.Println(typeId)

}

func isOperator(bits *[]string, versionTot *int, expr *[]string) {

	lengthTypeId := (*bits)[0]
	fmt.Printf("THIS IS LENGTH TYPE: %v\n", lengthTypeId)

	*bits = (*bits)[1:]

	switch lengthTypeId {
	case "0":

		subPackageSizeBits := ""
		for i := 0; i < 15; i++ {

			subPackageSizeBits = subPackageSizeBits + (*bits)[i]

		}

		lenSubPackages2 := calcBin(subPackageSizeBits)
		//fmt.Println(*bits)

		fmt.Printf("TAMANHO BITES: %v\n", lenSubPackages2)

		bitsToProc := make([]string, 0)

		*bits = (*bits)[15:]

		bitsToProc = (*bits)[:lenSubPackages2]

		*bits = (*bits)[lenSubPackages2:]

		for len(bitsToProc) > 0 {

			decodePackage(&bitsToProc, versionTot, expr)
		}

		//if len(bitsExcesso) > 6 {
		//	*bits = bitsExcesso
		//}

	case "1":

		subPackageSizeBits := ""
		for i := 0; i < 11; i++ {

			subPackageSizeBits = subPackageSizeBits + (*bits)[i]

		}

		lenSubPackages := calcBin(subPackageSizeBits)

		//fmt.Println(*bits)

		fmt.Printf("NUMERO PACKAGES: %v\n", lenSubPackages)

		*bits = (*bits)[11:]

		for i := 0; i < lenSubPackages; i++ {
			//fmt.Println(*bits)
			decodePackage(bits, versionTot, expr)
		}

	}
	//*expr = append(*expr, string(typeId))
	*expr = append(*expr, ")")

}

func isNumber(bits *[]string) int {

	numTot, bitsUsed := 0, 0
	bitflag := true
	bitSize := len((*bits))
	var bitString string
	i := 0

	for i = 0; i < bitSize && bitflag; i++ {

		bitsUsed += 4

		if (*bits)[i] == "0" {
			bitflag = false
		}

		for j := 0; j < 4; j++ {
			i++
			bitString = bitString + (*bits)[i]
		}

	}

	(*bits) = (*bits)[i:]

	if len(*bits) < 6 {
		*bits = make([]string, 0)
	}

	numTot = calcBin(bitString)

	return numTot
}

func expNumber(base int, exp int) int {

	if exp == 0 {
		return 1
	}
	var res int = 1
	for i := 0; i < exp; i++ {
		res *= base
	}

	return res
}

func calcBin(bits string) int {

	finalVal := 0

	bitsSize := len(bits)
	j := 0
	for i := bitsSize - 1; i >= 0; i-- {

		bit, _ := strconv.Atoi(string(bits[j]))

		finalVal += (expNumber(2, i) * bit)

		j++
	}

	return finalVal

}

func pop(arr *[]string) string {

	val := "undefined"

	if len(*arr) > 0 {
		val = (*arr)[0]
		*arr = (*arr)[1:]
	}

	return val

}

func handleOp(op string, equacao *[]string) int {

	switch op {

	case "+":
		return soma(equacao)
	case "*":
		return multiCos(equacao)
	case "min":
		return getMinOp(equacao)
	case "max":
		return getMaxOp(equacao)
	case ">":
		return biggerThan(equacao)
	case "<":
		return smallerThan(equacao)
	case "=":
		return equalsNum(equacao)
	}

	return 0

}

func getMinOp(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)
	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	minVal := resEq[0]

	for _, val := range resEq {
		if val < minVal {
			minVal = val
		}
	}

	return minVal

}

func getMaxOp(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)
	maxVal := 0
	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	for _, val := range resEq {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal

}

func equalsNum(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)

	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	if resEq[0] == resEq[1] {
		return 1
	} else {
		return 0
	}

}

func smallerThan(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)

	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	if resEq[0] < resEq[1] {
		return 1
	} else {
		return 0
	}

}

func biggerThan(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)

	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	if resEq[0] > resEq[1] {
		return 1
	} else {
		return 0
	}

}

func multiCos(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)
	eqTot := 1
	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	for _, val := range resEq {
		eqTot *= val
	}

	return eqTot

}

func soma(equacao *[]string) int {

	val := pop(equacao)
	resEq := make([]int, 0)
	eqTot := 0
	for val != ")" {

		if val != "(" {
			strVal, err := strconv.Atoi(val)

			if err != nil {
				resEq = append(resEq, handleOp(val, equacao))
			} else {
				resEq = append(resEq, strVal)
			}
		}

		val = pop(equacao)

	}

	for _, val := range resEq {
		eqTot += val
	}

	return eqTot
}

var Input_pack string = "6051639005B56008C1D9BB3CC9DAD5BE97A4A9104700AE76E672DC95AAE91425EF6AD8BA5591C00F92073004AC0171007E0BC248BE0008645982B1CA680A7A0CC60096802723C94C265E5B9699E7E94D6070C016958F99AC015100760B45884600087C6E88B091C014959C83E740440209FC89C2896A50765A59CE299F3640D300827902547661964D2239180393AF92A8B28F4401BCC8ED52C01591D7E9D2591D7E9D273005A5D127C99802C095B044D5A19A73DC0E9C553004F000DE953588129E372008F2C0169FDB44FA6C9219803E00085C378891F00010E8FF1AE398803D1BE25C743005A6477801F59CC4FA1F3989F420C0149ED9CF006A000084C5386D1F4401F87310E313804D33B4095AFBED32ABF2CA28007DC9D3D713300524BCA940097CA8A4AF9F4C00F9B6D00088654867A7BC8BCA4829402F9D6895B2E4DF7E373189D9BE6BF86B200B7E3C68021331CD4AE6639A974232008E663C3FE00A4E0949124ED69087A848002749002151561F45B3007218C7A8FE600FC228D50B8C01097EEDD7001CF9DE5C0E62DEB089805330ED30CD3C0D3A3F367A40147E8023221F221531C9681100C717002100B36002A19809D15003900892601F950073630024805F400150D400A70028C00F5002C00252600698400A700326C0E44590039687B313BF669F35C9EF974396EF0A647533F2011B340151007637C46860200D43085712A7E4FE60086003E5234B5A56129C91FC93F1802F12EC01292BD754BCED27B92BD754BCED27B100264C4C40109D578CA600AC9AB5802B238E67495391D5CFC402E8B325C1E86F266F250B77ECC600BE006EE00085C7E8DF044001088E31420BCB08A003A72BF87D7A36C994CE76545030047801539F649BF4DEA52CBCA00B4EF3DE9B9CFEE379F14608"

var map_HexaToBin = []string{
	"0 = 0000",
	"1 = 0001",
	"2 = 0010",
	"3 = 0011",
	"4 = 0100",
	"5 = 0101",
	"6 = 0110",
	"7 = 0111",
	"8 = 1000",
	"9 = 1001",
	"A = 1010",
	"B = 1011",
	"C = 1100",
	"D = 1101",
	"E = 1110",
	"F = 1111"}
