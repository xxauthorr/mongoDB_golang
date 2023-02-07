package main

// package contains all the models for the mongo database

type InsertData struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Age           string `json:"age"`
	Qualification string `json:"qualification"`
}
	