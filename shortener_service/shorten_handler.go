package shortener_service

import (
	"net/http"
	"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"math/rand"
	"time"
)

type shortenRequest struct {
	Url string
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var request shortenRequest
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprintf(w, `{"error" : "Invalid request"}`)
	}

	fmt.Printf("URL : %s", request.Url);

	short_url := GenRandomStr()
	SaveURLToDB(request.Url, short_url)

	fmt.Fprintf(w, `{"url" : "%s"}`, short_url)
}

func GenRandomStr() string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, RANDOM_LENGTH);
	for i := 0; i < RANDOM_LENGTH; i++ {
		result[i] =  RANDOM_SOURCE[rand.Intn(len(RANDOM_SOURCE))]
	}

	return string(result)
}

func SaveURLToDB(longURL string, shortURL string) int64 {
	db, err := sql.Open(DRIVER_NAME, DB_NAME)
	ErrToPanic(err)

	insert, err := db.Prepare("insert into urls(long_url, short_url) values(?, ?)")
	ErrToPanic(err)

	result, err := insert.Exec(longURL, shortURL)
	ErrToPanic(err)

	id, err := result.LastInsertId()
	ErrToPanic(err)

	return id
}
