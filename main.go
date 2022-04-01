package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Get Method
type Game struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Platform    string   `json:"platform"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Released    string   `json:"released"`
	Category    []string `json:"category"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/games", getGames).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":1234", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getGames(w http.ResponseWriter, r *http.Request) {
	games, err := fetchGameData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(&games)
	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Write(jsonResponse)
}

func fetchGameData() ([]Game, error) {
	jsonFile, err := os.Open("DB.json")
	if err != nil {
		return []Game{}, err
	}
	defer jsonFile.Close()

	games := []Game{}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []Game{}, err
	}

	json.Unmarshal(byteValue, &games)
	return games, nil
}
