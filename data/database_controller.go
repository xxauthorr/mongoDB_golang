package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID            string    `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName     string    `bson:"first_name" json:"first_name"`
	LastName      string    `bson:"last_name" json:"last_name"`
	Email         string    `bson:"email" json:"email"`
	Age           string    `bson:"age" json:"age"`
	Qualification string    `bson:"qualification" json:"qualification"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}

// insert an entry into the collection
func (l *LogEntry) Insert(entry LogEntry) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("students").Collection("entries")

	res, err := collection.InsertOne(ctx, LogEntry{
		FirstName:     entry.FirstName,
		LastName:      entry.LastName,
		Age:           entry.Age,
		Qualification: entry.Qualification,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into students:", err)
		return nil, err
	}
	var id struct {
		ObjectId any `json:"object_id"`
	}
	id.ObjectId = res.InsertedID

	return id, nil
}

// to get all entries from the collection students
func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("students").Collection("entries")
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "first_name", Value: -1}})
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []*LogEntry
	for cursor.Next(ctx) {
		var item LogEntry
		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error in decoding log into slice:", err)
			return nil, err
		} else {
			students = append(students, &item)
		}
	}

	return students, nil
}

// to get single entry using id
func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database("students").Collection("entries")

	var entry LogEntry
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error in changing format of object id:", err)
		return nil, err
	}
	result := collection.FindOne(ctx, bson.M{"_id": docID})
	err = result.Decode(&entry)
	if err != nil {
		log.Println("Error in decoding the result", err)
		return nil, err
	}
	return &entry, nil
}

// to drop the collection
func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("students").Collection("entries")
	if err := collection.Drop(ctx); err != nil {
		// log.Println("error drop collection :", err)
		return err
	}
	return nil
}

// to delete one entry
func (l *LogEntry) DeleteOne(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	collection := client.Database("students").Collection("entries")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error in changing format of object id:", err)
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		log.Println("Error in deleting data:", err)
		return err
	}
	return nil
}

// to update the entry
func (l *LogEntry) Update(data LogEntry) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := client.Database("students").Collection("entries")

	docID, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		log.Println("Error in changing format of object id:", err)
		return nil, err
	}
	result, err := collection.UpdateOne(ctx,
		bson.M{"_id": docID},
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "first_name", Value: data.FirstName},
					{Key: "last_name", Value: data.LastName},
					{Key: "email", Value: data.Email},
					{Key: "age", Value: data.Age},
					{Key: "Qualification", Value: data.Qualification},
					{Key: "updated_at", Value: time.Now()},
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}
	return result, nil
}
