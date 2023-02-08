package main

import (
	"context"
	"fmt"
	"log"
	"my_app/data"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort = "3000"
	// mongoURL = "mongodb://localhost:27017"
	mongoURL = "saavy.b3bokeg.mongodb.net/?retryWrites=true&w=majority"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {

	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {

		log.Println(err)
	}

	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	log.Printf("Starting service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connectToMongo() (*mongo.Client, error) {
	username := os.Getenv("username")
	password := os.Getenv("password")
	// create connection options
	// username := "admin"
	// password := "roots199970"
	dbURL := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, mongoURL)
	clientOptions := options.Client().ApplyURI(dbURL)
	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}
	return c, nil
}
