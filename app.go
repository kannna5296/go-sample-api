package main

import (
    "fmt"
    "net/http"
    "log"
	"time"
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

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "sample"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)
	fmt.Println(dsn)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()
  r := chi.NewRouter()
  r.Get("/handler", handler)
  http.ListenAndServe(":8080", r)
}