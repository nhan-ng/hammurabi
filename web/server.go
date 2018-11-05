package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nhan-ng/hammurabi/web/api"
)

const serverPort = ":9000"

var games = map[string]api.Game{
	"ham_1": api.Game{
		Name:        "ham_1",
		Description: "First game",
	},
	"ham_2": api.Game{
		Name:        "ham_2",
		Description: "Second game",
	},
}

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/games", gamesHandler)
	apiRouter.HandleFunc("/games/{name}", gameHandler)

	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/build/"))))
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./frontend/build"))))

	log.Printf("Server is running on port %s\n", serverPort)
	log.Fatalln(http.ListenAndServe(serverPort, router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Hammurabi!")
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	g := make([]api.Game, 0, len(games))
	for _, val := range games {
		g = append(g, val)
	}

	// Create a json payload from that
	json.NewEncoder(w).Encode(&g)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if g, ok := games[vars["name"]]; ok {
		json.NewEncoder(w).Encode(&g)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
