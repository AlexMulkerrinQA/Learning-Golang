package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
	"sync"
)

type safeCommands struct {
	commands map[int]string
	mux sync.Mutex
}
func (c *safeCommands) addCommand(id int, cmd string) {
	c.mux.Lock()
	c.commands[id] = cmd
	c.mux.Unlock()
}

type UpdateStruc struct {
	Agents []Agent
	Changes []mapChange
}

func gameLoadHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("client.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(content))
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := json.Marshal(game.Terrain)
	fmt.Fprintf(w, "%s", string(result))
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	playerID := 0
	found := false
	for i:=0; i<playerNum && found==false; i++ {
		if playerAssigned[i] == false {
			playerID = i
			found = true
			playerAssigned[i] = true
			game.Agents[i].Active = true
		}
	}
	if found {
		fmt.Printf("Player %v joined.\n", playerID)
		fmt.Fprintf(w, "%v", playerID)
	} else {
		fmt.Fprintf(w, "sorry no free slots :(")
	}
	
}

func leaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Player leaving.")
	givenid := r.FormValue("id")
	id, err := strconv.Atoi(givenid)
	if err != nil || id<0 ||id>=playerNum {
		//not a valid player id
		return
	}
	playerAssigned[id] = false
	game.Agents[id].Active = false
	fmt.Printf("Player %v left.\n", id)
	
	
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	upd := UpdateStruc{game.Agents, game.Changes}
	result, _ := json.Marshal(upd)
	fmt.Fprintf(w, "%s", string(result))
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	direc := r.FormValue("direc")
	givenid := r.FormValue("id")
	id, err := strconv.Atoi(givenid)
	if err != nil || id<0 ||id>=playerNum {
		//not a valid player id
		return
	}
	cp := &commands
	cp.addCommand(id, direc)
	fmt.Fprintf(w, "Sent command: %v", direc)
}

var (	
	game GameState
	playerAssigned map[int]bool
	commands safeCommands
	playerNum int = 13
	mapSize int = 13
)
func main() {
	game = Game(playerNum, mapSize)
	playerAssigned = make(map[int]bool)
	commands = safeCommands{commands:make(map[int]string)}
	for i:=0; i<playerNum; i++ {
		playerAssigned[i] = false
	}
	
	http.HandleFunc("/", gameLoadHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/leave", leaveHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/move", moveHandler)
	go http.ListenAndServe(":80", nil)
	
	tick := time.Tick(100*time.Millisecond)
	for  {
		<-tick // wait for clock tick
		// safely handle commands
		gp := &game
		cp := &commands
		cp.mux.Lock()
		gp.update(cp.commands)
		for i:=0; i<playerNum; i++ {
			cp.commands[i] = ""
		}
		cp.mux.Unlock()
		
		// prune old map changes
		var newChanges []mapChange
		for i:=0; i<len(game.Changes); i++ {
			game.Changes[i].ttl--
			if game.Changes[i].ttl>0 {
				newChanges = append(newChanges, game.Changes[i])
			}
		}
		game.Changes = newChanges
		
		
	}
}