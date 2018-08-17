package shortener_service

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
)

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["url"]

	longURL := GetLongURL(shortURL)

	w.Header().Set("Location", longURL)
	w.WriteHeader(301)
}

func GetLongURL(shortURL string) string {
	db, err := sql.Open(DRIVER_NAME, DB_NAME)
	if err != nil {
		return ""
	}

	get, err := db.Prepare("select long_url from urls where short_url = ?")
	if err != nil {
		return ""
	}

	rows, err := get.Query(shortURL)
	if err != nil {
		return ""
	}

	var longURL string
	rows.Next()
	rows.Scan(&longURL)

	return longURL;
}