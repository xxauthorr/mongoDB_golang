package main

// package contains all the models for the mongo database

type StudentData struct {
	Id            string `json:"id,omitempty"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Age           string `json:"age"`
	Qualification string `json:"qualification"`
}
