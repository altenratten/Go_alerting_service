package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var gstorage = make(map[string]float64)
var cstorage = make(map[string]int64)

func Handlers(w http.ResponseWriter, r *http.Request) {
	// get the value for the greeting wildcard.
	t := r.PathValue("type")
	n := r.PathValue("name")
	v := r.PathValue("value")

	switch t {
	case `gauge`:

		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		gstorage[n] = s

		w.Header().Set("Content-Type", "text/plain; charset=utf-8") 
		w.WriteHeader(http.StatusOK)

	case `counter`:
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8") 
		w.WriteHeader(http.StatusOK)

		cstorage[n] = s

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func storageHandlers(w http.ResponseWriter, r *http.Request){
	body := `------gstorage------\r\n`
    w.Write([]byte(body))
	jsonString, err := json.Marshal(gstorage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(jsonString))
	body2 := `------cstorage------\r\n`
    w.Write([]byte(body2))
	jsonString2, err := json.Marshal(cstorage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(jsonString2))
}

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/update/{type}/{name}/{value}", Handlers)
	mux.HandleFunc("/storage/", storageHandlers)

	//mux.Handle(`POST /update/{type}/{name}/{value}`, Handlers.UpdateHandlers(storage))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
