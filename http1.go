package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/slideshow/{id}/{width}/{height}", SlideShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

type SlideShowData struct {
	Id     string `json:"id"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Url    string `json:"url"`
}

func SlideShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payload := SlideShowData{vars["id"], vars["width"], vars["height"], "http://s3.amazon.fake/" + vars["id"]}

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
