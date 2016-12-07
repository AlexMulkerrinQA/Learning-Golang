package main

import (
	"fmt"
	"math/rand"
)

type GameState struct {
	Terrain [][]int
	Agents []Agent
}

func (this *GameState) update(cmd [10]string) {
	for i:=0; i<9; i++ {
		switch cmd[i] {
			case "left" :
				this.Agents[i].X -= 1
				if this.Agents[i].X<0 {
					this.Agents[i].X = 8
				}
			case "right" :
				this.Agents[i].X += 1
				if this.Agents[i].X>8 {
					this.Agents[i].X = 0
				}
			case "up" :
			this.Agents[i].Y -= 1
				if this.Agents[i].Y<0 {
					this.Agents[i].Y = 8
				}
			case "down" :
				this.Agents[i].Y += 1
				if this.Agents[i].Y>8 {
					this.Agents[i].Y = 0
				}
		}
	}

	/* for i:=0; i<9; i++ {
		x := this.Agents[i].X
		x = x + rand.Intn(3) -1
		if x<0 {
			x=0
		} else if x>=9 {
			x=8
		}
		y := this.Agents[i].Y
		y = y + rand.Intn(3) -1
		if y<0 {
			y=0
		} else if y>=9 {
			y=8
		}
		this.Agents[i].X = x
		this.Agents[i].Y = y
	} */
}

type Agent struct {
	X int
	Y int
}

func Game() GameState{
	var terrain = createMap(9, 9)
	printGrid(terrain)
	var agents []Agent
	for i:=0; i<9; i++ {
		a := Agent{rand.Intn(9), rand.Intn(9)}
		agents = append(agents, a)
	}
	return GameState{terrain, agents}
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