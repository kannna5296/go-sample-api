package main

import (
    "fmt"
	"os"
    "net/http"
    "log"
	"os"
	"strconv"
	"github.com/jmoiron/sqlx"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *sqlx.DB
)

type Book struct {
	ID          int       `db:"id"`
	Title string    `db:"account_name"`
	Author    string    `db:"passhash"`
	ReleaseDate time.Time `db:"created_at"`
}

const baseMessage = "HELLO WORLD"

func handler(w http.ResponseWriter, r *http.Request) {

	query := "SELECT * FROM `book`"
	var books []Book
	err = db.Select(&books, query)

	name := r.FormValue("name") //クエリパラメータ取得
	if name == "" {
		fmt.Fprintf(w, baseMessage) //ResponseWriterを渡してレスポンスを返す
		return
	}
	
	fmt.Fprintf(w, "%s, %s", baseMessage, name) //ResponseWriterを渡してレスポンスを返す
}

func main() {
  fmt.Print(os.Getppid())
  r := chi.NewRouter()
  r.Get("/handler", handler)
  http.ListenAndServe(":8080", r)
}