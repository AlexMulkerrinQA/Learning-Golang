package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)


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
	for i:=0; i<10 && found==false; i++ {
		if playerAssigned[i] == false {
			playerID = i
			found = true
			playerAssigned[i] = true
		}
	}
	if found {
		fmt.Printf("Player %v joined.\n", playerID)
		fmt.Fprintf(w, "%v", playerID)
	} else {
		fmt.Fprintf(w, "sorry no free slots :(")
	}
	
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := json.Marshal(game.Agents)
	fmt.Fprintf(w, "%s", string(result))
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	direc := r.FormValue("direc")
	givenid := r.FormValue("id")
	id, err := strconv.Atoi(givenid)
	if err != nil {
		//not a valid player id
		return
	}
	commands[id] = direc
	fmt.Fprintf(w, "TODO: moving %s Direction:%s", r.URL.Path[1:], direc)
}

var (	
	game GameState
	playerAssigned map[int]bool
	commands [10]string
)
func main() {
	game = Game()
	playerAssigned = make(map[int]bool)
	for i:=0; i<10; i++ {
		playerAssigned[i] = false
	}
	commands[0] = "left"
	
	http.HandleFunc("/", gameLoadHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/move", moveHandler)
	go http.ListenAndServe(":80", nil)
	
	tick := time.Tick(100*time.Millisecond)
	for  {
		// wait for clock tick
		<-tick
		gp :=&game
		gp.update(commands)
		for i:=0; i<10; i++ {
			commands[i] = ""
		}
		//fmt.Println("update!")
	}
}