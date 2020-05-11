package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Instantiate Mongo client
var client *mongo.Client

// Author model
type Author struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Stories   []Story
}

// Story model
type Story struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}

func createStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func main() {
	fmt.Println("Starting application...")

	// Connect to MongoDB
	fmt.Println("Connecting MongoDB")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// Instantiate router.
	r := chi.NewRouter()

	// Run app
	port := "5000"
	http.ListenAndServe(port, r)
}
