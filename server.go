package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

func getTime(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	timeWithTimestamp := currentTime.Format(time.RFC3339)
	data := map[string]string{
		"time": timeWithTimestamp,
	}
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if req.Method == "GET" {
		fmt.Fprintf(w, string(b))
	} else {
		fmt.Fprintf(w, "Sorry, only GET request is supported.")
	}

func main() {
	http.HandleFunc("/time", getTime)
	fmt.Printf("Starting server on ")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}

