package main

import (
	"errors"
	"log"
	"my_app/data"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (app *Config) InsertData(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	// read json
	var requestPayload InsertData

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("error in reading json"))
		return
	}
	event := data.LogEntry{
		FirstName:     requestPayload.FirstName,
		LastName:      requestPayload.LastName,
		Email:         requestPayload.Email,
		Age:           requestPayload.Age,
		Qualification: requestPayload.Qualification,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "Data uploaded successfully"

	app.writeJSON(w, http.StatusAccepted, jsonResponse)

}
