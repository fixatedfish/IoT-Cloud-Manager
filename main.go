package main

import (
	"IoTManager/controllers"
	"IoTManager/middleware"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"goji.io"
	"goji.io/pat"
	"log"
	"net/http"
	"os"
	"time"
)

var PORT = "8080"

func main() {
	setupLogging()

	mux := goji.NewMux()
	mux.Use(middleware.Logger)
	mux.HandleFunc(pat.Post("/config/update/:deviceId"), controllers.UpdateIot)
	mux.HandleFunc(pat.Get("/health"), controllers.Health)

	err := http.ListenAndServe("0.0.0.0:"+PORT, mux)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Listening on localhost", PORT)
}

func setupLogging() {
	// log to custom file
	LOGFILE := "/var/log/iot-webserver/webserver_" + time.DateOnly + ".log"
	// open log file
	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
