package main

import (
	"fmt"
	"math/rand"
)

type mapChange struct {
	X int
	Y int
	Value int
	ttl int
}

type GameState struct {
	Terrain [][]int
	Agents []Agent
	Changes []mapChange
}
func (this *GameState) update(cmd map[int]string) {
	for i:=0; i<len(this.Agents); i++ {
		if this.Agents[i].Active {
			switch cmd[i] {
				case "left" :
					this.tryMove(i, -1, 0)
				case "right" :
					this.tryMove(i, 1, 0)
				case "up" :
					this.tryMove(i, 0, -1)
				case "down" :
					this.tryMove(i, 0, 1)
				case "pick" :
					this.tryPick(i)
					
				case "drop" :
					this.tryDrop(i)
			}
		}
	}
}
func (this *GameState) tryPick(a int) bool {
	x := this.Agents[a].X
	y := this.Agents[a].Y
	if this.Agents[a].Carrying || this.Terrain[y][x]<2 {
		fmt.Printf("player %v cant pick", a)
		return false
	}
	//this.Terrain[y][x] -=1
	this.changeMap(-1, x, y)
	this.Agents[a].Carrying = true
	fmt.Printf("player %v picking", a)
	return true
}
func (this *GameState) tryDrop(a int) bool {
	x := this.Agents[a].X
	y := this.Agents[a].Y
	if this.Agents[a].Carrying == false{
		return false
	}
	//this.Terrain[y][x] +=1
	this.changeMap(1, x, y)
	this.Agents[a].Carrying = false
	return true
}
func (this *GameState) tryMove(a, dx, dy int) bool{
	nx := this.Agents[a].X + dx
	ny := this.Agents[a].Y + dy
	// check bounds
	if (nx<0 || nx>=len(this.Terrain) || ny<0 || ny>=len(this.Terrain)) {
		return false
	}
	// collide with other agents
	for i := 0; i<len(this.Agents); i++ {
		if a != i {
			if nx == this.Agents[i].X && ny == this.Agents[i].Y {
				return false
			}
		}
	}
	// collide with unpassable tiles
	if this.Terrain[ny][nx] == 0 {
		return false
	}
	
	// move is valid!
	this.Agents[a].X = nx
	this.Agents[a].Y = ny
	return true
}


func (this *GameState) changeMap(delta, x, y int) {
	this.Terrain[y][x] += delta
	this.Changes = append(this.Changes, mapChange{x,y,this.Terrain[y][x],60})
}

type Agent struct {
	X int
	Y int
	Active bool
	Carrying bool
}

func Game(playerNum, mapSize int) GameState{
	var terrain = createMap(mapSize, mapSize)
	var changes []mapChange
	var agents []Agent
	for i:=0; i<playerNum; i++ {
		a := Agent{rand.Intn(mapSize), rand.Intn(mapSize), false, false}
		agents = append(agents, a)
	}
	
	fmt.Printf("running %vx%v map with %v players.\n", mapSize, mapSize, playerNum)
	return GameState{terrain, agents, changes}
}

func createMap(width, height int) [][]int {
	var result [][]int
	for j := 0; j < height; j++ {
		var s []int
		for i := 0; i < width; i++ {
			r := rand.Intn(4)
			s = append(s, r)
		}
		result = append(result, s)
	}
	return result
}