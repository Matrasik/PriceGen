package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const PORT = ":8888"

var MapCurr = make(map[string][]int)
var MapDate = make(map[string][]time.Time)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", DefaultHandler)
	router.HandleFunc("/{name}", WriterHandlerCurr).Methods("GET")
	server := &http.Server{
		Addr:         PORT,
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		return
	}
}

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s\n		", "<3 <3")
}

func WriterHandlerCurr(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(req)
	json.NewEncoder(w).Encode(struct {
		Name string      `json:"name"`
		Curr []int       `json:"curr"`
		Date []time.Time `json:"date"`
	}{vars["name"], MapCurr[vars["name"]], MapDate[vars["name"]]})
}
