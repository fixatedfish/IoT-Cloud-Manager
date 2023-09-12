package controllers

import (
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Heath Check: OK")
}
