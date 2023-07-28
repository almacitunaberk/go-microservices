package main

import (
	"logger-service/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// Read the JSON into a variable
	var requestPayload jsonPayload

	_ = app.readJSON(w, r, &requestPayload)

	// Insert Data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "logged",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}