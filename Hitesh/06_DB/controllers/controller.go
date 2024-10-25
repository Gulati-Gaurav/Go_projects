package controller

import (
	"context"
	model "dbProj/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbConnectionString = "mongodb+srv://gauravgulati:iKesa3DK0LPTS4vW@cluster0.hlzlx0c.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// init is a method in golang which runs at the start of the program and only once
// With init(), you can run your code before the main() function and after the initialization of global variables.
func init() {

	// client options
	clientOption := options.Client().ApplyURI(mongodbConnectionString)

	// connect to db
	client, err := mongo.Connect(context.TODO(), clientOption)
	// Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	// should use todo in the initialisation and background while making a request

	if err != nil {
		logFatal(err)
	}

	fmt.Println("MongoDB Connection established")

	collection = client.Database(dbName).Collection((collectionName))

	fmt.Println("Collection instance ready")
}

// insert in db
func insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	logFatal(err)
	fmt.Println("Inserted in DB with ID = ", inserted.InsertedID)
}

func updateMovie(movieId string) {
	objId, err := primitive.ObjectIDFromHex(movieId)
	logFatal(err)

	// have to provide 2 things, one what to filter on and what to update
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	logFatal(err)

	fmt.Println("Modified count = ", res.ModifiedCount)
}

func deleteMovie(movieId string) {
	objId, err := primitive.ObjectIDFromHex(movieId)
	logFatal(err)

	filter := bson.M{"_id": objId}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	logFatal(err)
	fmt.Println("Number of records delete = ", deleted.DeletedCount)
}

func deleteAllMovies(movieId string) {
	// see the diff between bson.M and bson.D

	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}) // {} for empty i.e. filter all
	logFatal(err)
	fmt.Println("Number of records delete = ", deleted.DeletedCount)
}

func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	logFatal(err)
	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		logFatal(err)
		movies = append(movies, movie)
	}
	defer curr.Close(context.Background())
	return movies
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// actual controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
