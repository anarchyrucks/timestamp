package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Time struct {
	Unix    *int64  `json:"unix"`
	Natural *string `json:"natural"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/{time}", TimeHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/index.html")
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	str := mux.Vars(r)["time"]
	w.WriteHeader(http.StatusOK)

	if timestmp, err := strconv.Atoi(str); err == nil {
		t := time.Unix(int64(timestmp), 0)

		timestring := t.Format("January 02, 2006")
		timestamp := int64(timestmp)

		json.NewEncoder(w).Encode(Time{
			&timestamp,
			&timestring,
		})
	} else {
		t, err := time.Parse("January 02, 2006", str)
		if err != nil {
			json.NewEncoder(w).Encode(Time{
				nil,
				nil,
			})
		} else {
			timestamp := t.Unix()
			timestring := str

			json.NewEncoder(w).Encode(Time{
				&timestamp,
				&timestring,
			})
		}
	}
}
