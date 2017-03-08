package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Time struct {
	Unix    int64  `json:"unix"`
	Natural string `json:"natural"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/{time}", TimeHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/index.html")
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	str := mux.Vars(r)["time"]
	w.WriteHeader(http.StatusOK)
	if timestamp, err := strconv.Atoi(str); err == nil {
		t := time.Unix(int64(timestamp), 0)
		timestring := t.Format("January 02, 2006")
		json.NewEncoder(w).Encode(Time{
			int64(timestamp),
			timestring,
		})
	} else {
		t, err := time.Parse("January 02, 2006", str)
		if err != nil {
			fmt.Fprintf(w, "Invalid input")
			os.Exit(0)
		}
		timestamp := t.Unix()
		timestring := str
		json.NewEncoder(w).Encode(Time{
			timestamp,
			timestring,
		})
	}
}
