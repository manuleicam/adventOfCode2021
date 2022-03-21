package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	vertices []*Vertice
}

type Vertice struct {
	val      string
	nextVert []*Vertice
}

func (g *Graph) print() {

	for _, ver := range g.vertices {

		fmt.Printf("From: %v To: ", ver.val)
		for _, val := range ver.nextVert {
			fmt.Printf("%v ", val.val)
		}
		fmt.Println()
	}
	fmt.Println()

}

func (g *Graph) addElement(from, to string) {

	g.addVertice(from)
	g.addVertice(to)

	fromVer := g.getVertice(from)
	toVer := g.getVertice(to)

	if !(fromVer == nil || toVer == nil) {
		g.addNext(&fromVer.nextVert, toVer)
		g.addNext(&toVer.nextVert, fromVer)
	}

}

func (g *Graph) addNext(from *[]*Vertice, to *Vertice) {

	if !contains(*from, to.val) {
		*from = append(*from, to)
	}

}

func (g *Graph) addVertice(val string) {

	if !contains(g.vertices, val) {
		g.vertices = append(g.vertices, &Vertice{val: val})
	}

}

func (g *Graph) getVertice(vertice string) *Vertice {

	for _, ver := range g.vertices {

		if ver.val == vertice {
			return ver
		}

	}

	return nil

}

func contains(vertices []*Vertice, vertice string) bool {

	for _, val := range vertices {
		if val.val == vertice {
			return true
		}
	}

	return false
}

func countPath(fromName string, from Vertice, to string, visited []string, path []string, smallCaveTwice bool) int {

	if from.val == to {
		return 0
	}

	totalPath := 0

	path = append(path, from.val)
	if strings.ToLower(from.val) == from.val {
		visited = append(visited, from.val)
	}

	for _, adjVert := range from.nextVert {
		if adjVert.val == to {

			path = append(path, adjVert.val)
			//fmt.Printf("%s ", visited)
			fmt.Println(path)
			totalPath++

		} else if !checkVisited(visited, adjVert.val) {
			totalPath += countPath(fromName, *adjVert, to, visited, path, smallCaveTwice)

		} else if adjVert.val != fromName {
			if !smallCaveTwice {
				//break
				//totalPath += countPath(*adjVert, to, visited, path, smallCaveTwice)
				totalPath += countPath(fromName, *adjVert, to, visited, path, true)

			}
		}

	}

	return totalPath

}

func checkVisited(visited []string, ver string) bool {

	for _, verVisited := range visited {
		if verVisited == ver {
			return true
		}
	}

	return false
}

func (g *Graph) findAllPath(from, to string) int {

	if from == to {
		return 0
	}

	fromVert := g.getVertice(from)
	//toVert := g.getVertice(to)

	var visited = []string{from}
	var path = []string{}

	totalPath := countPath(from, *fromVert, to, visited, path, false)

	return totalPath

}

func main() {

	g := Graph{}

	for _, path := range InputList {
		splitInput := strings.Split(path, "-")
		g.addElement(splitInput[0], splitInput[1])
	}

	totalPath := g.findAllPath("start", "end")

	g.print()
	fmt.Println(totalPath)

}

var InputList = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end"}

var InputList_2 = []string{
	"dc-end",
	"HN-start",
	"start-kj",
	"dc-start",
	"dc-HN",
	"LN-dc",
	"HN-end",
	"kj-sa",
	"kj-HN",
	"kj-dc"}

var InputList_3 = []string{
	"kc-qy",
	"qy-FN",
	"kc-ZP",
	"end-FN",
	"li-ZP",
	"yc-start",
	"end-qy",
	"yc-ZP",
	"wx-ZP",
	"qy-li",
	"yc-li",
	"yc-wx",
	"kc-FN",
	"FN-li",
	"li-wx",
	"kc-wx",
	"ZP-start",
	"li-kc",
	"qy-nv",
	"ZP-qy",
	"nv-xr",
	"wx-start",
	"end-nv",
	"kc-nv",
	"nv-XQ"}
