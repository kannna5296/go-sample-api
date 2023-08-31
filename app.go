package main

import (
    "fmt"
    "net/http"
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
  fmt.Printf("HTTPサーバ立ち上げ")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}