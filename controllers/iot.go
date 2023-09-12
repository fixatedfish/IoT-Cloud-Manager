package controllers

import (
	"IoTManager/models"
	"database/sql"
	"encoding/json"
	"goji.io/pat"
	"io"
	"log"
	"net/http"
)

func UpdateIot(w http.ResponseWriter, r *http.Request) {
	deviceId := pat.Param(r, "deviceId")
	db, err := sql.Open("sqlite3", "iot.db")
	previousDp, _ := models.GetDataPoints(db, &deviceId)
	dp := new(models.DataPoints)
	body := make([]byte, r.ContentLength)
	_, err = io.ReadFull(r.Body, body)
	err = json.Unmarshal(body, dp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid Request Body"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = dp.UpdateConfig(db, deviceId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataPointsResult, err := models.GetDataPoints(db, &deviceId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updated, _ := json.Marshal(dataPointsResult)
	w.Write(updated)

	if previousDp["waterLevelLow"]["value"] == false && (*dp)["waterLevelLow"]["value"] == true {
		log.Println("Triggering email...")
	}
}
