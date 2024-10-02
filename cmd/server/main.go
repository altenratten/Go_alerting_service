package main

import (
	"fmt"
	"net/http"
)

func mainPage (res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
    body += "Header ===============\r\n"
    for k, v := range req.Header {
        body += fmt.Sprintf("%s: %v\r\n", k, v)
    }
    body += "Query parameters ===============\r\n"
    for k, v := range req.URL.Query() {
        body += fmt.Sprintf("%s: %v\r\n", k, v)
    }
    res.Write([]byte(body))
}

func apiGauge(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("You in apiGauge!"))
}

func apiCounter (res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("You in apiCounter!"))
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))

	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/update/gauge/", apiGauge)
	mux.HandleFunc("/update/counter/", apiCounter)

    //mux.Handle(`POST /update/{type}/{name}/{value}`, handlers.UpdateHandler(storage))





	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}