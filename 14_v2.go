package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	polyCod := make([]string, 0)
	polyMap := make(map[string]string, 0)
	cache := make(map[string]map[string]int, 0)
	charAppearance := make(map[string]int, 0)

	for _, char := range Input_list {

		polyCod = append(polyCod, string(char))

	}

	buildMapPoly(&polyMap, Input_list_Map)

	//for i := 0; i < 10; i++ {

	//parte_1
	fmt.Println(polyCod)
	polymerization(polyMap, polyCod, &charAppearance, &cache, 40)

	//parte_2
	//polymerization(polyMap, polyCod, &charAppearance, &cache, 40,40)

	//}

	//fmt.Println(polyMap)
	//fmt.Println(polyCod)
	//fmt.Println()
	fmt.Println(charAppearance)
	fmt.Println(difMostAndLessCommon(charAppearance))
	//fmt.Println(cache)
}

func difMostAndLessCommon(charAppearance map[string]int) int {

	mostPop := 0
	lessPop := -1

	for _, val := range charAppearance {

		if lessPop == -1 || lessPop > val {
			lessPop = val
		}

		if mostPop < val {
			mostPop = val
		}

	}

	return mostPop - lessPop

}

func polymerization(polyMap map[string]string, polyCod []string, charAppearance *map[string]int, cache *map[string]map[string]int, stepsToFinish int) {

	stringSize := len(polyCod)

	i := 0

	for i = 0; i < stringSize-1; i++ {

		resCache, existCache := (*cache)[polyCod[i]+polyCod[i+1]+strconv.Itoa(stepsToFinish)]

		if existCache {
			actLetterAppearanceCacheBased(polyCod[i], polyCod[i+1], stepsToFinish, charAppearance, cache)
			//atualizar as aparencias das letras com base na cache
			fmt.Println("CACHE_V2")
			fmt.Println(resCache)
		} else {

			aux := make(map[string]int, 0)
			aux[polyCod[i]] = 1

			(*cache)[polyCod[i]+polyCod[i+1]+strconv.Itoa(stepsToFinish)] = aux

			val, exist := (*charAppearance)[(polyCod)[i]]

			if exist {
				(*charAppearance)[(polyCod)[i]] = val + 1
			} else {
				(*charAppearance)[(polyCod)[i]] = 1
			}

			nCharsGen := buildPolyPair(polyMap, (polyCod)[i], (polyCod)[i+1], charAppearance, cache, stepsToFinish-1)

			actCache(polyCod[i], (polyCod)[i+1], stepsToFinish, nCharsGen, cache)

		}

		//fmt.Println(*charAppearance)

	}

	(*charAppearance)[polyCod[i]]++

}

func actLetterAppearanceCacheBased(char1, char2 string, stepsToFinish int, charAppearance *map[string]int, cache *map[string]map[string]int) {

	cacheVal, existCache := (*cache)[char1+char2+strconv.Itoa(stepsToFinish)]

	if existCache {

		for charKey, val := range cacheVal {

			charVal, exist := (*charAppearance)[charKey]

			if exist {
				(*charAppearance)[charKey] = charVal + val
			} else {
				(*charAppearance)[charKey] = val
			}

		}

	}

}

func buildPolyPair(polyMap map[string]string, char1, char2 string, charAppearance *map[string]int, cache *map[string]map[string]int, stepsToFinish int) map[string]int {

	charAppear := make(map[string]int)

	if stepsToFinish < 0 {
		return charAppear
	}

	resCache, existCache := (*cache)[char1+char2+strconv.Itoa(stepsToFinish)]
	//fmt.Println(*cache)

	if existCache {
		//atualizar as aparencias das letras com base na cache
		actLetterAppearanceCacheBased(char1, char2, stepsToFinish, charAppearance, cache)

		//fmt.Println("CACHE", char1, char2, stepsToFinish)
		//fmt.Println(resCache)

		return resCache
	} else {

		resPair, existPair := polyMap[char1+char2]

		aux := make(map[string]int, 0)
		aux[resPair] = 1

		(*cache)[char1+char2+strconv.Itoa(stepsToFinish)] = aux

		if existPair {

			val, exist := (*charAppearance)[resPair]

			if exist {
				(*charAppearance)[resPair] = val + 1
			} else {
				(*charAppearance)[resPair] = 1
			}

			charAppear[resPair] = 1

			charAppearChar1 := buildPolyPair(polyMap, char1, resPair, charAppearance, cache, stepsToFinish-1)
			actCache(char1, char2, stepsToFinish, charAppearChar1, cache)
			charAppear = joinResRec(charAppearChar1, charAppear)

			charAppearChar2 := buildPolyPair(polyMap, resPair, char2, charAppearance, cache, stepsToFinish-1)
			actCache(char1, char2, stepsToFinish, charAppearChar2, cache)
			charAppear = joinResRec(charAppearChar2, charAppear)

		}

	}

	return charAppear

}

func joinResRec(charAppearance1 map[string]int, charAppearance2 map[string]int) map[string]int {

	for key, val := range charAppearance1 {

		charVal, exist := charAppearance2[key]

		if exist {
			charAppearance2[key] = charVal + val
		} else {
			charAppearance2[key] = val
		}

	}

	return charAppearance2

}

func actCache(char1, resPair string, stepsToFinish int, charAppearance map[string]int, cache *map[string]map[string]int) {

	for char, val := range charAppearance {

		//fmt.Println((*cache)[char1+resPair+strconv.Itoa(stepsToFinish)])

		cacheVal, exist := (*cache)[char1+resPair+strconv.Itoa(stepsToFinish)][char]

		if exist {
			(*cache)[char1+resPair+strconv.Itoa(stepsToFinish)][char] = cacheVal + val
		} else {

			aux := (*cache)[char1+resPair+strconv.Itoa(stepsToFinish)]
			aux[char] = val

			(*cache)[char1+resPair+strconv.Itoa(stepsToFinish)] = aux
		}

	}

}

func buildMapPoly(polyMap *map[string]string, inputMap []string) {

	for _, oneMap := range inputMap {

		splitInput := strings.Split(oneMap, " -> ")

		(*polyMap)[splitInput[0]] = splitInput[1]

	}

}

var Input_list_teste = "NNCB"

var Input_list = "CPSSSFCFOFVFNVPKBFVN"

var Input_list_Map = []string{"NV -> V",
	"CF -> O",
	"BB -> F",
	"SB -> H",
	"KF -> O",
	"SP -> H",
	"CS -> V",
	"VF -> F",
	"PC -> H",
	"PH -> H",
	"SF -> F",
	"CP -> B",
	"BC -> H",
	"PB -> V",
	"KO -> B",
	"CV -> S",
	"ON -> B",
	"PV -> F",
	"OO -> B",
	"VV -> B",
	"NO -> B",
	"SH -> N",
	"FC -> C",
	"VO -> B",
	"NN -> C",
	"HH -> S",
	"CK -> C",
	"PF -> N",
	"SN -> K",
	"OK -> F",
	"FH -> S",
	"BP -> K",
	"HO -> K",
	"FB -> P",
	"HC -> N",
	"FP -> P",
	"NC -> H",
	"PK -> O",
	"BV -> P",
	"HK -> S",
	"PP -> N",
	"VC -> K",
	"CH -> C",
	"KS -> V",
	"KB -> V",
	"FN -> P",
	"BS -> O",
	"PS -> N",
	"NS -> B",
	"PN -> N",
	"NP -> H",
	"CB -> S",
	"SV -> O",
	"OC -> H",
	"BO -> C",
	"HN -> N",
	"HP -> N",
	"OF -> H",
	"FS -> S",
	"KV -> P",
	"HV -> S",
	"VS -> P",
	"BH -> N",
	"CC -> V",
	"VN -> H",
	"NF -> B",
	"NK -> N",
	"CN -> F",
	"FV -> P",
	"NH -> S",
	"OV -> H",
	"KN -> F",
	"SO -> H",
	"OP -> N",
	"KC -> P",
	"HB -> B",
	"BN -> V",
	"VP -> N",
	"HS -> S",
	"VK -> C",
	"VH -> H",
	"OS -> C",
	"FO -> B",
	"NB -> P",
	"KP -> V",
	"SS -> O",
	"BK -> F",
	"SK -> N",
	"HF -> O",
	"PO -> F",
	"OH -> B",
	"KK -> O",
	"FK -> S",
	"VB -> V",
	"OB -> C",
	"KH -> H",
	"SC -> F",
	"FF -> H",
	"CO -> V",
	"BF -> H"}

var Input_list_Map_test = []string{
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C"}
