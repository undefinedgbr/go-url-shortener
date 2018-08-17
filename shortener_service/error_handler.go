package shortener_service

import (
	"net/http"
	"fmt"
)

func ErrToPanic(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ErrToJson(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Fprintf(w, `{"error" : "%s"`, err)
	}
}