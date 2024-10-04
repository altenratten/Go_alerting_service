package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var gstorage = make(map[string]float64)
var cstorage = make(map[string]int64)

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`http.StatusNotFound`))
	// body := fmt.Sprintf("Method: %s\r\n", req.Method)
	// body += "Header ===============\r\n"
	// for k, v := range req.Header {
	//     body += fmt.Sprintf("%s: %v\r\n", k, v)
	// }
	// body += "Query parameters ===============\r\n"
	// body += `<a href="http://127.0.0.1:8080/update/gggg/rrr/21">link text</a>`
	// for k, v := range req.URL.Query() {
	//     body += fmt.Sprintf("%s: %v\r\n", k, v)
	// }
	// res.Write([]byte(body))
}

func Handlers(w http.ResponseWriter, r *http.Request) {
	// get the value for the greeting wildcard.
	t := r.PathValue("type")
	n := r.PathValue("name")
	v := r.PathValue("value")

	switch t {
	case `gauge`:

		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			w.Write([]byte(`http.StatusBadRequest`))
		}
		gstorage[n] = s

		jsonString, _ := json.Marshal(gstorage)

		w.Write([]byte(jsonString))

	case `counter`:
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			w.Write([]byte(`http.StatusBadRequest`))
		}

		cstorage[n] = s

		// jsonString, _ := json.Marshal(gstorage)

		// w.Write([]byte(jsonString))

	default:
		w.Write([]byte(`http.StatusNotFound`))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/update/{type}/{name}/{value}", Handlers)

	//mux.Handle(`POST /update/{type}/{name}/{value}`, Handlers.UpdateHandlers(storage))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
