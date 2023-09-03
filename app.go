package main

import (
    "fmt"
    "net/http"
    "log"
	"github.com/go-chi/chi/v5"
)

const baseMessage = "HELLO WORLD"

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //クエリパラメータ取得
	if name == "" {
		fmt.Fprintf(w, baseMessage) //ResponseWriterを渡してレスポンスを返す
		return
	}
	
	fmt.Fprintf(w, "%s, %s", baseMessage, name) //ResponseWriterを渡してレスポンスを返す
}

func main() {
  log.Fatal("Application started.")
  r := chi.NewRouter()
  r.Get("/handler", handler)
  http.ListenAndServe(":8080", r)
}