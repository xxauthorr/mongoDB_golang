package main

import (
	"errors"
	"my_app/data"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (app *Config) InsertOne(w http.ResponseWriter, r *http.Request) {
	// read json
	var requestPayload StudentData

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

	id, err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "Request successfully completed"
	jsonResponse.Data = id

	app.writeJSON(w, http.StatusAccepted, jsonResponse)

}

func (app *Config) DeleteOne(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Id string `json:"id"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("error in reading json"))
		return
	}
	err = app.Models.LogEntry.DeleteOne(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "Request successfully completed"
	app.writeJSON(w, http.StatusAccepted, jsonResponse)
}

func (app *Config) GetOne(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Id string `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("error in reading json"))
		return
	}
	data, err := app.Models.LogEntry.GetOne(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "Request successfully completed"
	jsonResponse.Data = data
	app.writeJSON(w, http.StatusAccepted, jsonResponse)
}

func (app *Config) Update(w http.ResponseWriter, r *http.Request) {

	var requestPayload StudentData

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("error in reading json"))
		return
	}
	event := data.LogEntry{
		ID:            requestPayload.Id,
		FirstName:     requestPayload.FirstName,
		LastName:      requestPayload.LastName,
		Email:         requestPayload.Email,
		Age:           requestPayload.Age,
		Qualification: requestPayload.Qualification,
	}

	_, err = app.Models.LogEntry.Update(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var data struct {
		Id any `json:"updated_id"`
	}
	data.Id = requestPayload.Id

	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "Request successfully completed"
	jsonResponse.Data = data
	app.writeJSON(w, http.StatusAccepted, jsonResponse)

}

func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := app.Models.LogEntry.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	if data == nil {
		jsonResponse.Message = "collection is empty"
		app.readJSON(w, r, jsonResponse)
		return
	}
	jsonResponse.Message = "data successfully collected"
	jsonResponse.Data = data

	app.writeJSON(w, http.StatusAccepted, jsonResponse)
}

func (app *Config) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	if err := app.Models.LogEntry.DropCollection(); err != nil {
		app.errorJSON(w, err)
		return
	}
	var jsonResponse jsonResponse
	jsonResponse.Error = false
	jsonResponse.Message = "collection deleted successfully"
	app.writeJSON(w, http.StatusAccepted, jsonResponse)

}
