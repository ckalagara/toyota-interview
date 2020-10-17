package main

import (
	"log"
)

//Grid :
//[
//["A","B","C","E"],
//["S","F","C","S"],
//["A","D","E","E"]
//]
//
//
//
//“ABCCEDX” : false
//“ECCEDAS” : true
//“DEES” : true
//
//Write a function which takes in a string and this grid and return true if the word can be formed from this grid. Up down left right. Only visit cell once
func main() {

	grid := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "C", "S"},
		{"A", "D", "E", "E"},
	}

	log.Printf(" “ABCCEDX” : false | %v", canForm("ABCCEDX", grid))
	log.Printf(" “ECCEDAS” : true | %v", canForm("ECCEDAS", grid))
	log.Printf(" “DEES” : true | %v", canForm("DEES", grid))
}

func canForm(w string, grid [][]string) bool {

	for _, position := range positions(string(w[0]), grid) {
		if ok := tryAPosition(position[0], position[1], w, grid); ok {
			return true
		}
	}

	return false
}

func tryAPosition(i, j int, w string, grid [][]string) bool {
	log.Printf("------------------------------------------checking for %v with %v's grid position: %v, %v", w, string(w[0]), i, j)
	var ok bool
	for li, l := range w {

		if li == 0 {
			continue
		}
		//log.Printf("-------------looking for ------- %v", string(l))
		in, jn, okn := isANeightbour(i, j, string(l), grid)
		if okn {
			i = in
			j = jn
			ok = okn
			continue
		}
		return false
	}
	return ok
}

func isANeightbour(i, j int, next string, grid [][]string) (int, int, bool) {

	if i != 0 && j != 0 {
		// check left neighbour
		ni, nj := i, j-1

		if nj >= 0 {
			nv := grid[ni][nj]

			if next == nv {
				log.Printf("left | %v | %v,%v > %v", next, ni, nj, nv)
				return ni, nj, true
			}
		}
	}

	// check up
	ni, nj := i-1, j
	if i == 0 {
		ni = len(grid) - 1
	}
	nv := grid[ni][nj]

	if next == nv {
		log.Printf("  up | %v | %v,%v > %v", next, ni, nj, nv)
		return ni, nj, true
	}

	// check down
	ni, nj = i+1, j
	if i == len(grid)-1 {
		ni = 0
	}
	nv = grid[ni][nj]

	if next == nv {
		log.Printf("down | %v | %v,%v > %v", next, ni, nj, nv)
		return ni, nj, true
	}

	if j != len(grid[len(grid)-1])-1 {
		// check right
		if len(grid[i]) > j+1 {
			ni, nj := i, j+1
			nv := grid[ni][nj]

			if next == nv {
				log.Printf("right | %v | %v,%v > %v", next, ni, nj, nv)
				return ni, nj, true
			}
		}
	}

	return 0, 0, false
}

func positions(s string, grid [][]string) [][]int {

	var positions [][]int

	for i, set := range grid {
		for j, l := range set {
			if l == s {
				positions = append(positions, []int{i, j})

			}
		}
	}

	return positions
}
