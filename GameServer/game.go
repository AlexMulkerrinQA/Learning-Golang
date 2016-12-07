package main

import (
	"fmt"
	"math/rand"
)

func Game() [][]int{
	var terrain = createMap(9, 9)
	printGrid(terrain)
	return terrain
}

func createMap(width, height int) [][]int {
	var result [][]int
	for j := 0; j < height; j++ {
		var s []int
		for i := 0; i < width; i++ {
			r := rand.Intn(9) + 1
			s = append(s, r)
		}
		result = append(result, s)
	}
	return result
}

func printGrid( value [][]int ) {
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			fmt.Print("+-")
		}
		fmt.Print("+\n")
		for j:=0; j<9; j++ {
			fmt.Print("|")
			fmt.Print(value[j][i])
		}
		fmt.Print("|\n")
	}
	for j:=0; j<9; j++ {
		fmt.Print("+-")
	}
	fmt.Print("+\n")
}