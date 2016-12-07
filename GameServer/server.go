package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
	result, _ := json.Marshal(terrain)
	fmt.Fprintf(w, "%s", string(result))
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: player %s", r.URL.Path[1:])
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: update %s", r.URL.Path[1:])
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	var direc string = r.FormValue("direc")
	fmt.Fprintf(w, "TODO: moving %s Direction:%s", r.URL.Path[1:], direc)
}

var terrain [][]int

func main() {
	terrain = Game()
	


	http.HandleFunc("/", gameLoadHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/move", moveHandler)
	http.ListenAndServe(":8080", nil)
}