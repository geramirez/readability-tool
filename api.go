package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"scorer"
)

// serveStats collects stats from post request and returns stats in json format
func serveStats(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Methods", "POST")
	body, _ := ioutil.ReadAll(req.Body)
	stats := scorer.GetStats(string(body))
	statsJSON, _ := json.Marshal(stats)
	io.WriteString(rw, string(statsJSON))
}

func init() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api", serveStats)
	http.ListenAndServe(":8000", nil)
}
