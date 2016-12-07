package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
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
	result, _ := json.Marshal(game.Agents)
	fmt.Fprintf(w, "%s", string(result))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: update %s", r.URL.Path[1:])
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	var direc string = r.FormValue("direc")
	fmt.Fprintf(w, "TODO: moving %s Direction:%s", r.URL.Path[1:], direc)
}

var game GameState

func main() {
	game = Game()
	
	http.HandleFunc("/", gameLoadHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/move", moveHandler)
	go http.ListenAndServe(":80", nil)
	
	tick := time.Tick(time.Second)
	for  {
		// wait for clock tick
		<-tick
		gp :=&game
		gp.update()
		fmt.Println("update!")
	}
}