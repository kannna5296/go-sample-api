package main

import (
	"log"
    "net/http"
	"encoding/json"
)

const baseMessage = "HELLO WORLD"

func handler(w http.ResponseWriter, r *http.Request) {
	// name := r.FormValue("name") //クエリパラメータ取得
	u := User{ Id : 1, Name : "Thome", Email : "thome@example.com"}
    json, err := json.Marshal(u)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write(json)
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error ListenAndServe : ", err)
    }
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}